package handlers

import (
	"groupe-tracker/internal/models"
	"groupe-tracker/internal/services"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func showDetailsPage(w http.ResponseWriter, r *http.Request, artists []models.Artist, brandID string, apiClient *services.APIClient) {
	id, err := strconv.Atoi(brandID)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, "Invalid brand ID")
		return
	}

	var artist models.Artist
	for _, a := range artists {
		if a.ID == id {
			artist = a
			break
		}
	}

	if artist.ID == 0 {
		ErrorHandler(w, http.StatusNotFound, "Artist not found")
		return
	}

	relationData, err := apiClient.FetchRelationData(artist.Relations)
	if err != nil {
		log.Printf("Failed to fetch relation data: %v", err)
		ErrorHandler(w, http.StatusInternalServerError, "Failed to fetch relation data")
		return
	}

	tmpl, err := template.ParseFiles("internal/templates/detailsPage.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Failed to parse template")
		return
	}

	data := struct {
		Artist   models.Artist
		Relation models.RelationData
	}{
		Artist:   artist,
		Relation: relationData,
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Failed to execute template: %v", err)
		ErrorHandler(w, http.StatusInternalServerError, "Failed to render page")
	}
}
