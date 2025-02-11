package config

import (
	"os"
	"strconv"
)

type Config struct {
	MySQL *MySQLConfig
	JWT   *JWTConfig
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

type JWTConfig struct {
	JWTSecret string
}

func NewConfig() *Config {
	return &Config{MySQL: NewMySQLConfig(), JWT: NewJWTConfig()}
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

func NewJWTConfig() *JWTConfig {
	s := os.Getenv("SECRET_KEY")
	return &JWTConfig{JWTSecret: s}
}
