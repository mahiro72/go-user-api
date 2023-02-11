package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mahiro72/go-user-api/pkg/logger"
	"github.com/mahiro72/go-user-api/pkg/presenter/http/router"
	"github.com/mahiro72/go-user-api/pkg/registry"
)

func main() {
	repo, err := registry.NewRepository()
	if err != nil {
		panic(fmt.Sprintf("error: NewRepository: %v",err))
	}

	r := router.New(repo)

	srv := &http.Server{Addr: ":8080", Handler: r}
	srvCtx, srvStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// 猶予期間が30秒のグレースフルシャットダウンを開始する
		shutdownCtx, _ := context.WithTimeout(srvCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Log("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// グレースフルシャットダウンのトリガー
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			logger.Log(err.Error())
		}
		srvStopCtx()
	}()

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Log(err.Error())
	}

	// サーバーのcontextが終了するまで待機する
	<-srvCtx.Done()
}
