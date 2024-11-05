package util

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	location := location()
	godotenv.Load(fmt.Sprintf("%v/%v", location, ".env"))
}
