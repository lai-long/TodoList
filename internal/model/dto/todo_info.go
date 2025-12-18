package dto

type TodoList struct {
	Id      int    `json:"-"`
	Title   string `json:"title"`
	Context string `json:"context"`
	Status  string `json:"status"`
	EndDate string `json:"end_date"`
}
