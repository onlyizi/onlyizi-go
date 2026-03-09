package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv(files ...string) error {
	if len(files) == 0 {
		return godotenv.Load()
	}

	return godotenv.Load(files...)
}
