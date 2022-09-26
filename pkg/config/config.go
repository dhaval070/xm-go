package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DSN string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	cfg := Config{}

	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err)
	}

	viper.BindEnv("GOXM_DSN")
	cfg.DSN = viper.GetString("GOXM_DSN")

	return cfg
}
