package main

import (
	"github.com/hyeyoom/go-web-app-boilerplate/provider"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect migration")
	}

	migrator := provider.Migrator{Database: db}
	migrator.Migrate()
}
