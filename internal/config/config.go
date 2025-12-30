package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisAddr     string
	RedisPassword string

	JWTSecret string
}

func Load() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	// 允许从环境变量读取
	viper.AutomaticEnv()

	// 读取 .env（如果存在）
	if err := viper.ReadInConfig(); err != nil {
		log.Println("no .env file found, using environment variables")
	}

	return &Config{
		AppPort: viper.GetString("APP_PORT"),

		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),

		RedisAddr:     viper.GetString("REDIS_ADDR"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),

		JWTSecret: viper.GetString("JWT_SECRET"),
	}
}
