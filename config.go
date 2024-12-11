package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Database struct {
		Driver string `mapstructure:"driver"`
		File   string `mapstructure:"file"`
	} `mapstructure:"database"`
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	return config
}

func main() {
	config := LoadConfig()

	fmt.Printf("Server running on %s:%d\n", config.Server.Host, config.Server.Port)
	fmt.Printf("Database file: %s\n", config.Database.File)
}
