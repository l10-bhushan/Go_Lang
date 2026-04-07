package model

type TodoRequest struct {
	Title string `json:"title"`
}

type Todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Status    bool   `json:"status"`
	Timestamp string `json:"timestamp"`
}
