package main

import (
	"fmt"
	"groupe-tracker/config"
	"groupe-tracker/internal/handlers"
	"groupe-tracker/internal/services"
	"log"
	"net/http"
)

func main() {

	//setting up API client
	apiClient := services.NewAPIClient(config.BASE_URL)

	http.HandleFunc("/", handlers.HomeHandler(apiClient))

	log.Println(fmt.Sprintf("Server is running on port %s...", config.SERVER_PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.SERVER_PORT), nil))
}
