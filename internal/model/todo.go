package model

type TodoList struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Context string `json:"context"`
	Status  bool   `json:"status"`
	EndDate string `json:"end_date"`
}
