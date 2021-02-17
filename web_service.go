package main

import (
	"belajar-golang/helper"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func palindromService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := []struct {
		Palindrom string
		Count     string
	}{
		{"1 10", "9"},
		{"99 100", "1"},
		{"21 31", "1"},
	}

	dataJson, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		r.ParseForm()
		inputPalindrom := r.FormValue("palindrom")
		transform := strconv.Itoa(helper.Palindrom(inputPalindrom))
		w.Write([]byte("total number palindrom is " + transform))
	case "GET":
		w.Write(dataJson)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

}

func main() {
	http.HandleFunc("/palindrom", palindromService)

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
