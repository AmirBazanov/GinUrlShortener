package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	pgUri = ""
)

func main() {
	var migrationsPath, migrationsTable string
	flag.StringVar(&migrationsPath, "migrationsPath", "./migrations", "Path to the migrations folder")
	flag.StringVar(&migrationsTable, "migrationsTable", "migrations_test", "Name of the migrations table")
	flag.Parse()
	if migrationsPath == "" {
		panic("StoragePath and MigrationsPath are required")
	}
	m, err := migrate.New("file://"+migrationsPath, pgUri)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No change")
			return
		}
		panic(err)
	}
	fmt.Println("Migrations successfully migrated")
}
