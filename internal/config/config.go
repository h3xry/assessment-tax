package config

import "github.com/spf13/viper"

type ENV struct {
	Port          string
	AdminUsername string
	AdminPassword string
	DatabaseURL   string
}

func NewConfig() *ENV {
	viper.AutomaticEnv()
	cfg := ENV{
		Port:          viper.GetString("PORT"),
		AdminUsername: viper.GetString("ADMIN_USERNAME"),
		AdminPassword: viper.GetString("ADMIN_PASSWORD"),
		DatabaseURL:   viper.GetString("DATABASE_URL"),
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return &cfg
}
