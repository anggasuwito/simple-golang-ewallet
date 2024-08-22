package config

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	val = new(Configuration)
)

type Configuration struct {
	DBMaster                  *gorm.DB
	HttpHost                  string
	HttpPort                  string
	AppVersion                string
	AccessTokenSecret         string
	AccessTokenExpireDuration string
}

func SetConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[config] error when loading .env file")
	}

	dbMaster, err := getDatabase(dbConfig{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
		sslMode:  os.Getenv("DB_SSL"),
		timezone: os.Getenv("DB_TIMEZONE"),
	})
	if err != nil {
		log.Fatal("[config] failed connecting database")
	}

	val.DBMaster = dbMaster
	val.HttpHost = os.Getenv("HTTP_HOST")
	val.HttpPort = os.Getenv("HTTP_PORT")
	val.AppVersion = os.Getenv("APP_VERSION")
	val.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	val.AccessTokenExpireDuration = os.Getenv("ACCESS_TOKEN_EXPIRE_DURATION")
}

func GetConfig() *Configuration {
	return val
}
