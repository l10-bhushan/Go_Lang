package model

type Result struct {
	Status bool   `json:"status"`
	Data   string `json:"data"`
}

type CreationSuccess struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Todo `json:"data"`
}

type TodoResult struct {
	Status bool   `json:"status"`
	Data   []Todo `json:"data"`
}
