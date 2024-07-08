package vars

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"os"
)

var (
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_HOST string
	DB_PORT string
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading the env file. %s", err.Error())
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
}
