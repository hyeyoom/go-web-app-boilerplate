package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
)

func main() {
	password := "1234abcd"
	member := domain.NewMember("abc", password)
	fmt.Println(member)
	fmt.Println(member.IsCredentialVerified(password))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
