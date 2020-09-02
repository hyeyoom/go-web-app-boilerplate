package provider

import (
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
	"gorm.io/gorm"
)

type Migrator struct {
	Database *gorm.DB
}

func (sm *Migrator) Migrate() {
	err := sm.Database.AutoMigrate(domain.GetModels()...)
	if err != nil {
		log = GetLogger()
		log.Fatal("Error occurred during migration.", err)
		panic("Error occurred during migration.")
	}
}
