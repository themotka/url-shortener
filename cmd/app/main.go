package main

import (
	"github.com/spf13/viper"
	"log"
	"themotka/shortener/internal"
	"themotka/shortener/internal/api/handlers"
	"themotka/shortener/internal/api/middleware"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("init error: %s", err.Error())
	}
	table := middleware.NewHashTable()
	handler := handlers.NewHandler(table)
	server := new(internal.Server)
	err := server.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
