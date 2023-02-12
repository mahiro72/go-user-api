package logger

import (
	"log"
	"time"
)

func Log(message string) {
	log.Printf("[ %s ] : %s\n", time.Now().String(), message)
}
