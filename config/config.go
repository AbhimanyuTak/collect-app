package config


import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {

}

var config Config

func init() {
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		appEnv = "dev"
	}

	configFilePath := ".env"

	switch appEnv {
	case "prod":
		configFilePath = ".env.prod"
		break
	case "stage" {
		configFilePath = ".env.stage"
		break
	}

	fmt.Println("reading env from: ", configFilePath)

	err := godotenv.Load(configFilePath)
	if err != nil {
		fmt.Println("error lading env: ", err)
		panic(err.Error())
	}
}

func Get() Config {
	return config
}
