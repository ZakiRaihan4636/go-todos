package routes

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/zakiraihan4636/go-todos/controllers"
)

func TodoRoutes(e *echo.Echo, db *sql.DB) {
	e.POST("/todos", func(c echo.Context) error { return controllers.CreateTodo(c, db) })
	e.GET("/todos", func(c echo.Context) error { return controllers.GetTodos(c, db) })
	e.PATCH("/todos/:id", func(c echo.Context) error { return controllers.UpdateTodo(c, db) })
	e.DELETE("/todos/:id", func(c echo.Context) error { return controllers.DeleteTodo(c, db) })
	e.PATCH("/todos/:id/check", func(c echo.Context) error { return controllers.CheckTodo(c, db) })
	e.PATCH("/todos/:id/uncheck", func(c echo.Context) error { return controllers.UncheckTodo(c, db) })
}
