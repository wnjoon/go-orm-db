package util

import (
	"github.com/joho/godotenv"
)

func Env() {
	err := godotenv.Load(".env")
	if err != nil {
		HandleErr(err, ErrLoadEnv)
	}
}
