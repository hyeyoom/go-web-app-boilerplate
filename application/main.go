package main

import (
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
	"github.com/hyeyoom/go-web-app-boilerplate/provider"
	"github.com/hyeyoom/go-web-app-boilerplate/provider/persistence"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
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

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		log := provider.GetLogger()
		log.Info("ㅇㅅㅇ?!")
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
