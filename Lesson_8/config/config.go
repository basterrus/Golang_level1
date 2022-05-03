package config

import "os"

type Configuration struct {
	DbUrl      string
	DbPort     string
	DbUser     string
	DbPassword string
}

func ReadConfig() *Configuration {
	return &Configuration{
		DbUrl:      os.Getenv("DB_URL"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
}
