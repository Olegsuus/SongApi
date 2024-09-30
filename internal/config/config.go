package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DataBaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"DBName"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DataBaseConfig `yaml:"database"`
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// Переопределяем конфиг из переменных окружения
	cfg.Database.User = os.Getenv("POSTGRES_USER")
	cfg.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Database.DBName = os.Getenv("POSTGRES_DB")
	cfg.Database.Host = "db" // Используем имя сервиса в Docker Compose

	return &cfg
}
