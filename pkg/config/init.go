package config

import "fmt"

func init() {
	if err := initMySQL(); err != nil {
		panic(
			fmt.Sprintf("error: config.init failed: %v", err),
		)
	}
}
