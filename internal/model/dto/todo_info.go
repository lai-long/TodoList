package dto

type TodoList struct {
	UserName string `json:"username" `
	Title    string `json:"title"`
	Context  string `json:"context"`
	Status   string `json:"status"`
	EndDate  string `json:"end_date"`
}
