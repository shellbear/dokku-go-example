package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/shellbear/dokku-go-example/models"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	postgresURL := os.Getenv("DATABASE_URL")
	if postgresURL == "" {
		log.Fatalln("missing DATABASE_URL ENV variable.")
	}

	// Connect to database.
	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect to database:", err)
	}

	// Create Todo model migrations.
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalln("failed to run migrations:", err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It works!")
	})

	e.GET("/todos", GetAllTodos(db))
	e.POST("/todos", CreateNewTodo(db))
	e.GET("/todos/:id", GetOneTodo(db))
	e.PUT("/todos/:id", UpdateOneTodo(db))

	e.Logger.Fatal(e.Start(":"+port))
}
