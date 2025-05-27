package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zakiraihan4636/go-todos/database"

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

	err := db.Ping()

	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Create todo
	e.POST("/todos", func(c echo.Context) error {
		var req CreateRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		if req.Title == "" || req.Description == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Title and description are required"})
		}

		result, err := db.Exec("INSERT INTO todos (title, description, done) VALUES (?, ?, 0)", req.Title, req.Description)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert todo"})
		}

		lastID, _ := result.LastInsertId()

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Todo created successfully",
			"id":      lastID,
		})
	})

	// Get all todos
	e.GET("/todos", func(c echo.Context) error {
		rows, err := db.Query("SELECT id, title, description, done FROM todos")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to query todos"})
		}
		defer rows.Close()

		todos := []TodoResponse{}
		for rows.Next() {
			var todo TodoResponse
			if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan todo"})
			}
			todos = append(todos, todo)
		}

		return c.JSON(http.StatusOK, todos)
	})

	// Update todo
	e.PATCH("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		var req UpdateRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		if req.Title == "" || req.Description == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Title and description are required"})
		}

		// Check todo exists
		var exists int
		err := db.QueryRow("SELECT COUNT(*) FROM todos WHERE id = ?", id).Scan(&exists)
		if err != nil || exists == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
		}

		var doneInt int
		if req.Done == 1 {
			doneInt = 1
		}
		_, err = db.Exec("UPDATE todos SET title = ?, description = ?, done = ? WHERE id = ?", req.Title, req.Description, doneInt, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update todo"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Todo updated successfully"})
	})

	// Delete todo
	e.DELETE("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		result, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete todo"})
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Todo deleted successfully"})
	})

	// Mark todo as done
	e.PATCH("/todos/:id/check", func(c echo.Context) error {
		id := c.Param("id")

		_, err = db.Exec("UPDATE todos SET done = 1 WHERE id = ?", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to mark todo as done"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Todo marked as done"})
	})

	// Mark todo as undone
	e.PATCH("/todos/:id/uncheck", func(c echo.Context) error {
		id := c.Param("id")

		_, err = db.Exec("UPDATE todos SET done = 0 WHERE id = ?", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to mark todo as done"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Todo marked as undone"})
	})

	// server running on port 8080
	e.Start(":8080")
}
