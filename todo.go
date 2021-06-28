package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/shellbear/dokku-go-example/models"
)


// Create a new Todo.
func CreateNewTodo(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var todo models.Todo
		if err := c.Bind(&todo); err != nil {
			return err
		}

		log.Println("Creating a new Todo.")

		if err := db.Create(&todo).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, &todo)
	}
}

// List all Todos.
func GetAllTodos(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var todos []*models.Todo

		log.Println("Fetching all Todos")

		// Find all Todos in database.
		if err := db.Find(&todos).Error; err != nil {
			return err
		}

		// Return all the Todos as JSON.
		return c.JSON(http.StatusOK, todos)
	}
}

// Get Todo from ID.
func GetOneTodo(db *gorm.DB) func (c echo.Context) error {
	return func (c echo.Context) error {
		// Parse the ID url parameter.
		var id uint
		if err := echo.PathParamsBinder(c).
			Uint("id", &id).
			BindError(); err != nil {
			return err
		}

		log.Println("Fetching one Todo:", id)

		// Find the Todo from ID.
		var todo models.Todo
		if err := db.Find(&todo, "id = ?", id).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, &todo)
	}
}


// Update Todo from ID.
func UpdateOneTodo(db *gorm.DB) func (c echo.Context) error {
	return func(c echo.Context) error {
		var input models.Todo

		// Get JSON input.
		if err := c.Bind(&input); err != nil {
			return err
		}

		// Parse the ID url parameter.
		var id uint
		if err := echo.PathParamsBinder(c).
			Uint("id", &id).
			BindError(); err != nil {
			return err
		}

		log.Println("Updating one Todo:", id)

		input.ID = id

		// Update the Todo in database.
		if err := db.Updates(&input).Error; err != nil {
			return err
		}

		// Get the updated Todo.
		if err := db.Find(&input, "id = ?", id).Error; err != nil {
			return err
		}

		// Return the Todo as JSON.
		return c.JSON(http.StatusOK, &input)
	}
}

