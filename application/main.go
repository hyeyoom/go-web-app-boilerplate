package main

import (
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
	"github.com/hyeyoom/go-web-app-boilerplate/provider"
	"github.com/hyeyoom/go-web-app-boilerplate/provider/persistence"
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

	m := domain.NewMember("mail@mail.com", "1234abcd")

	sf := persistence.NewStorage(db)
	mr := sf.NewMemberRepository()
	mr.Create(m)

}
