package config

import (
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DataBaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DataBaseConfig `yaml:"database"`
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.BindEnv("Server.Port", "SERVER_PORT")
	viper.BindEnv("Database.Driver", "DB_DRIVER")
	viper.BindEnv("Database.Host", "DB_HOST")
	viper.BindEnv("Database.Port", "DB_PORT")
	viper.BindEnv("Database.User", "DB_USER")
	viper.BindEnv("Database.Password", "DB_PASSWORD")
	viper.BindEnv("Database.DBName", "DB_NAME")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &cfg
}
