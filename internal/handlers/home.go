package handlers

import (
	"groupe-tracker/internal/models"
	"groupe-tracker/internal/services"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func HomeHandler(apiClient *services.APIClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			ErrorHandler(w, http.StatusNotFound, "Page not found")
			return
		}

		artists, err := apiClient.GetAllArtists()
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError, "Failed to fetch data from API")
			return
		}

		if brandID := r.URL.Query().Get("brand"); brandID != "" {
			showDetailsPage(w, r, artists, brandID, apiClient)
			return
		}

		// Handle search
		criteria := r.URL.Query().Get("criteria")
		query := r.URL.Query().Get("query")
		noResults := false

		if criteria != "" && query != "" {
			artists = filterArtists(artists, criteria, query)
			noResults = len(artists) == 0
		}

		if len(r.URL.Query()) > 0 && (criteria == "" || query == "") {
			ErrorHandler(w, http.StatusBadRequest, "Wrong query or criteria")
			return
		}

		tmpl, err := template.ParseFiles("internal/templates/index.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError, "Failed to parse template")
			return
		}

		fullURL := "http://" + r.Host

		data := struct {
			Artists   []models.Artist
			FullURL   string
			NoResults bool
			Query     string
		}{
			Artists:   artists,
			FullURL:   fullURL,
			NoResults: noResults,
			Query:     query,
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("Failed to execute template: %v", err)
			ErrorHandler(w, http.StatusInternalServerError, "Failed to render page")
		}
	}
}

func filterArtists(artists []models.Artist, criteria, query string) []models.Artist {
	var filtered []models.Artist
	query = strings.ToLower(query)

	for _, artist := range artists {
		var match bool
		switch criteria {
		case "all":
			//Search by all
			match = strings.Contains(strings.ToLower(artist.Name), query) || strings.Contains(strings.ToLower(artist.FirstAlbum), query) || strings.Contains(strings.ToLower(string(artist.CreationDate)), query)
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), query) {
					match = true
					break
				}
			}
		case "name":
			match = strings.Contains(strings.ToLower(artist.Name), query)
		case "members":
			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), query) {
					match = true
					break
				}
			}
		case "creationDate":
			match = strings.Contains(strings.ToLower(string(artist.CreationDate)), query)
		case "firstAlbum":
			match = strings.Contains(strings.ToLower(artist.FirstAlbum), query)
		}

		if match {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}
