package model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Addr    string `json:"addr"`
	Country string `json:"country"`
}
