package main

import (
	"groupe-tracker/internal/handlers"
	"groupe-tracker/internal/services"
	"log"
	"net/http"
)

func main() {
	apiClient := services.NewAPIClient("https://groupietrackers.herokuapp.com/api")

	http.HandleFunc("/", handlers.HomeHandler(apiClient))

	log.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
