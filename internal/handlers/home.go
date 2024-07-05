package handlers

import (
	"fmt"
	"groupe-tracker/internal/models"
	"groupe-tracker/internal/services"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

		var flatLocations []string

		for _, art := range artists {
			data := art.LocationData
			if locations, ok := data.([]interface{}); ok {
				for _, loc := range locations {
					if strLoc, ok := loc.(string); ok {
						exist := false
						for _, loc := range flatLocations {
							if loc == strLoc {
								exist = true
								break
							}
						}
						if !exist {
							flatLocations = append(flatLocations, strLoc)
						}
					}
				}
			}
		}

		if brandID := r.URL.Query().Get("brand"); brandID != "" {
			showDetailsPage(w, r, artists, brandID, apiClient)
			return
		}

		// Handle search
		criteria := r.URL.Query().Get("criteria")
		query := r.URL.Query().Get("query")
		noResults := false

		var filteredArts []models.Artist

		if criteria != "" && query != "" {
			filteredArts = filterArtists(artists, criteria, query)
			noResults = len(artists) == 0
		}

		tmpl, err := template.ParseFiles("internal/templates/index.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError, "Failed to parse template")
			return
		}

		fullURL := "http://" + r.Host

		data := struct {
			Artists      []models.Artist
			FilteredArts []models.Artist
			FullURL      string
			NoResults    bool
			Query        string
			Criteria     string
			Locations    []string
			filteredArts []models.Artist
		}{
			Artists:      artists,
			FullURL:      fullURL,
			NoResults:    noResults,
			Query:        query,
			Criteria:     criteria,
			Locations:    flatLocations,
			FilteredArts: filteredArts,
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
			match = strings.Contains(strings.ToLower(artist.Name), query) || strings.Contains(strings.ToLower(artist.FirstAlbum), query) || strings.Contains(strconv.Itoa(int(artist.CreationDate)), query)

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
			match = strings.Contains(strconv.Itoa(int(artist.CreationDate)), query)
			break
		case "firstAlbum":
			match = strings.Contains(strings.ToLower(artist.FirstAlbum), query)
			break
		case "location":
			fmt.Println("In location")
			if locations, ok := artist.LocationData.([]interface{}); ok {
				for _, loc := range locations {
					if strLoc, ok := loc.(string); ok {
						if strings.Contains(strings.ToLower(strLoc), query) {
							match = true
							break
						}
					}
				}

			}
		}

		if match {
			filtered = append(filtered, artist)
		}
	}

	return filtered
}
