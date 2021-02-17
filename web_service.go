package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getPalindrom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := []struct {
		Name string
		Age  int
	}{
		{"Richard Grayson", 24},
		{"Jason Todd", 23},
		{"Tim Drake", 22},
		{"Damian Wayne", 21},
	}

	dataJson, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		w.Write([]byte("post"))
	case "GET":
		w.Write(dataJson)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

}

func main() {
	http.HandleFunc("/", getPalindrom)

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
