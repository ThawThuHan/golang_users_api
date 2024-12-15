package configs

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	ServerIP   string
	ServerPort string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}
	var config Config
	config.ServerIP = os.Getenv("SERVER_IP")
	if config.ServerIP == "" {
		config.ServerIP = "localhost"
	}

	config.ServerPort = os.Getenv("SERVER_PORT")
	if config.ServerPort == "" {
		config.ServerPort = "8080"
	}

	config.DbHost = os.Getenv("DB_HOST")
	if config.DbHost == "" {
		config.DbHost = "localhost"
	}

	config.DbPort = os.Getenv("DB_PORT")
	if config.DbPort == "" {
		config.DbPort = "3306"
	}

	config.DbUser = os.Getenv("DB_USER")
	if config.DbUser == "" {
		config.DbUser = "root"
	}

	config.DbPassword = os.Getenv("DB_PASSWORD")
	if config.DbPassword == "" {
		config.DbPassword = "root"
	}

	config.DbName = os.Getenv("DB_NAME")
	if config.DbName == "" {
		config.DbName = "golang"
	}

	return &config, nil
}

func (c *Config) ConnectDB() (*gorm.DB, error) {
	dsn := c.DbUser + ":" + c.DbPassword + "@tcp(" + c.DbHost + ":" + c.DbPort + ")/" + c.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
