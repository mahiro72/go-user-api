package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/mahiro72/go-user-api/pkg/config"
	"github.com/mahiro72/go-user-api/pkg/logger"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.MySQL.DSN), &gorm.Config{})
	if err != nil {
		spentSec := 0
		for err != nil && spentSec <= 60 {
			db, err = gorm.Open(mysql.Open(config.MySQL.DSN), &gorm.Config{})
			logger.Log(
				fmt.Sprintf("failed: connect to db, time: %ds ...",spentSec),
			)
			spentSec += 5
			time.Sleep(time.Second * 5)
		}
	}
	logger.Log("success: connect to db")
	return db, nil
}
