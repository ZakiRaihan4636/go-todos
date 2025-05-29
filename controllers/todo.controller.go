package controllers

import (
	"database/sql"
	"net/http"

	"github.com/zakiraihan4636/go-todos/models"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context, db *sql.DB) error {
	var req models.CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.Title == "" || req.Description == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Title and description are required"})
	}

	result, err := db.Exec("INSERT INTO todos (title, description, done) VALUES (?, ?, 0)", req.Title, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create todo"})
	}

	lastID, _ := result.LastInsertId()
	return c.JSON(http.StatusCreated, echo.Map{"message": "Todo created", "id": lastID})
}

func GetTodos(c echo.Context, db *sql.DB) error {
	rows, err := db.Query("SELECT id, title, description, done FROM todos")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Query error"})
	}
	defer rows.Close()

	todos := []models.TodoResponse{}
	for rows.Next() {
		var todo models.TodoResponse
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Scan error"})
		}
		todos = append(todos, todo)
	}

	return c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c echo.Context, db *sql.DB) error {
	id := c.Param("id")
	var req models.UpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.Title == "" || req.Description == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Title and description required"})
	}

	var exists int
	err := db.QueryRow("SELECT COUNT(*) FROM todos WHERE id = ?", id).Scan(&exists)
	if err != nil || exists == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Todo not found"})
	}

	_, err = db.Exec("UPDATE todos SET title=?, description=?, done=? WHERE id=?", req.Title, req.Description, req.Done, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Update failed"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Todo updated"})
}

func DeleteTodo(c echo.Context, db *sql.DB) error {
	id := c.Param("id")
	result, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Delete failed"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Todo not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Todo deleted"})
}

func CheckTodo(c echo.Context, db *sql.DB) error {
	id := c.Param("id")
	_, err := db.Exec("UPDATE todos SET done = 1 WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Check failed"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Marked as done"})
}

func UncheckTodo(c echo.Context, db *sql.DB) error {
	id := c.Param("id")
	_, err := db.Exec("UPDATE todos SET done = 0 WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Uncheck failed"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Marked as undone"})
}
