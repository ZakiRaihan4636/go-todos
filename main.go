package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zakiraihan4636/go-todos/database"
	"github.com/zakiraihan4636/go-todos/routes"

	_ "github.com/go-sql-driver/mysql"
)

type CreateRequest struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        int    `json:"done"`
}

type TodoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        int    `json:"done"`
}
type DeleteRequest struct {
	Id int `json:"id"`
}

func main() {
	db := database.InitDb()
	defer db.Close()
	e := echo.New()

	routes.TodoRoutes(e, db)
	// server running on port 8080
	e.Start(":8080")
}
