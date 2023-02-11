package config

import (
	"fmt"
	"os"
)

func getEnv(envPath string) (string, error) {
	v := os.Getenv(envPath)
	if v == "" {
		return "", fmt.Errorf("%s is empty", envPath)
	}

	return v, nil
}
