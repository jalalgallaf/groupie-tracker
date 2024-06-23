package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, status int, message string) {
	tmpl, err := template.ParseFiles("internal/templates/error.html")
	if err != nil {
		log.Printf("Failed to parse error template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: message,
	}

	w.WriteHeader(status)
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Failed to execute error template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
