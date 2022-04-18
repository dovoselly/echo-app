package util

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Dotenv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
}
