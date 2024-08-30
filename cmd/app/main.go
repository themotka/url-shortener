package main

import (
	"flag"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"
	"log"
	"themotka/shortener/internal"
	"themotka/shortener/internal/api/handlers"
	"themotka/shortener/internal/api/middleware"
	"themotka/shortener/internal/database"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("init error: %s", err.Error())
	}

	isFlagged := flag.Bool("d", false, "Work with database")
	flag.Parse()
	db, err := database.NewTable(&database.Config{
		Host: "localhost",
		Port: "5432",
		User: "postgres",
		Pass: "123",
		Name: "postgres",
		Mode: "disable",
	})
	if err != nil {
		log.Fatal(err)
	}
	table := middleware.NewHashTable(db)
	router := handlers.NewRouter(table)
	if !*isFlagged {
		table.Repo = nil
	}
	server := new(internal.Server)
	err = server.Run(viper.GetString("port"), router.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
