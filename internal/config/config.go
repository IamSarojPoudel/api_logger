package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	Port          string
	DatabaseUrl   string
	DatabaseName  string
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS"),
		Port:          getEnv("PORT"),
		DatabaseUrl:   getEnv("DATABASE_URL"),
		DatabaseName:  getEnv("DATABASE_NAME"),
	}
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatalf("Environment variable %s not set", key)
	return ""
}
