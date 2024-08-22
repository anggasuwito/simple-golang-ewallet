package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host     string
	user     string
	password string
	dbName   string
	port     string
	sslMode  string
	timezone string
}

func getDatabase(config dbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.host,
		config.user,
		config.password,
		config.dbName,
		config.port,
		config.timezone,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
