package main

import (
	"net/http"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPatch:
	case http.MethodDelete:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tasks/", taskHandler)
	http.ListenAndServe(":8000", nil)
}
