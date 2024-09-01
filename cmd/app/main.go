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
	err = makeMigrations(db)
	defer db.Close()

	storage := url.NewStorage(*isFlagged, db)
	shortener := url.NewShortener()
	service := url.NewService(storage, shortener)
	handler := handlers.NewHandler(&service)

	server := new(internal.Server)
	err = server.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}

func makeMigrations(db *sql.DB) error {
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
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil && !errors.Is(migrate.ErrNoChange, err) {
		log.Fatal(err)
	}
	return err
}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
