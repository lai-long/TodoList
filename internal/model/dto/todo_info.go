package dto

type TodoList struct {
	Title   string `json:"title"`
	Context string `json:"context"`
	Status  string `json:"status"`
	EndDate string `json:"end_date"`
}
