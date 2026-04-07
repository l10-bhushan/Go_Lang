package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Name string `json:"name"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Backend bootcamp started 🚀")
}

func bootcampHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	d := Result{Name: "Bhushan"}
	jData, err := json.Marshal(d)
	if err != nil {
		fmt.Println("Failed to encode json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bootcamp", bootcampHandler)

	fmt.Println("Server running on PORT :8080")
	http.ListenAndServe(":8080", nil)
}
