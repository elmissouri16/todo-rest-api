package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string
}

func setupRouter() *gin.Engine {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Todo{})

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/todos", func(ctx *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		ctx.JSON(http.StatusOK, todos)
	})
	r.GET("/todos/add", func(ctx *gin.Context) {
		// get content from query
		content := ctx.Query("content")
		// create new todo
		todo := Todo{Content: content}
		// save to db
		db.Create(&todo)
		ctx.JSON(http.StatusOK, todo)

	})

	return r
}
func main() {
	r := setupRouter()
	r.Run(":9090")
}
