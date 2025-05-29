package models

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
