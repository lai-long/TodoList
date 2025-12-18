package dto

type TodoList struct {
	Id      int
	Title   string `json:"title"`
	Context string `json:"context"`
	Status  string `json:"status"`
	EndDate string `json:"end_date"`
}
