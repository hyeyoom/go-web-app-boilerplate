package provider

import (
	"fmt"
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
	"gorm.io/gorm"
)

type Migrator struct {
	Database *gorm.DB
}

func (sm *Migrator) Migrate() {
	err := sm.Database.AutoMigrate(domain.GetModels()...)
	if err != nil {
		// todo: logging here
		fmt.Println(err)
		panic("Unable to migrate models")
	}
}
