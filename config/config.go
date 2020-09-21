package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// AppConfig ...
type AppConfig struct {
	Debug bool
	Server serverConfig
	Db     dbConfig
}

type serverConfig struct {
	Port         int
	TimeoutRead  time.Duration
	TimeoutWrite time.Duration
}

type dbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

// ReadEnv ...
func ReadEnv() *AppConfig {
	viper.SetConfigName("config")

	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var appConfig AppConfig

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil
	}

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return nil
	}

	fmt.Println(appConfig)

	return &appConfig
}
