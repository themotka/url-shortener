package main

import (
	"database/sql"
	"errors"
	"flag"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"themotka/shortener/internal"
	"themotka/shortener/internal/api/handlers"
	"themotka/shortener/internal/url"
	postgres2 "themotka/shortener/internal/url/adapters/db/pg"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("init error: %s", err.Error())
	}
	isFlagged := flag.Bool("d", false, "Work with database")
	flag.Parse()
	db, err := postgres2.NewTable(&postgres2.Config{
		Host: viper.GetString("dbHost"),
		Port: viper.GetString("dbPort"),
		User: viper.GetString("dbUser"),
		Pass: viper.GetString("dbPass"),
		Name: viper.GetString("dbName"),
		Mode: viper.GetString("dbMode"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init success")
	makeMigrations(db)
	log.Println("Migrate success")
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatalf("db close error: %s", err.Error())
		}
	}(db)
	storage := url.NewStorage(*isFlagged, db)
	log.Println("Storage initialised")
	shortener := url.NewShortener()
	service := url.NewService(storage, shortener)
	handler := handlers.NewHandler(&service)

	server := new(internal.Server)
	log.Println("Server initialised")
	err = server.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}

func makeMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		viper.GetString("dbName"),
		driver,
	)
	if err != nil {
		log.Fatalf("Migrating error: %s", err)
	}
	err = m.Up()
	if err != nil && !errors.Is(migrate.ErrNoChange, err) {
		log.Fatalf("Migrating error: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("/internal/configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
