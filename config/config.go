package config

import (
	"os"
	"strconv"
)

type Config struct {
	MySQL *MySQLConfig
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

func NewConfig() *Config {
	return &Config{MySQL: NewMySQLConfig()}
}

func NewMySQLConfig() *MySQLConfig {
	u := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	h := os.Getenv("MYSQL_HOST")
	p, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	db := os.Getenv("MYSQL_DB")

	return &MySQLConfig{
		User:     u,
		Password: pw,
		Host:     h,
		Port:     p,
		DB:       db}
}
