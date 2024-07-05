# Groupe Tracker

Groupe Tracker is a web application that allows users to search and filter information about music artists. It provides an interactive interface with autocomplete search functionality across various artist attributes.

## Features

- Search artists by name, members, creation date, first album, and location
- Autocomplete suggestions for search queries
- Filter artists based on different criteria
- Responsive design for various screen sizes

## Project Structure

- `cmd/`
    - `server/`
        - `main.go`: The entry point of the application
- `config/`
    - `config.go`: Defines config values used throughout the project
- `internal/`
    - `handlers/`
        - `artist.go`: Handles artist-related HTTP requests
        - `error.go`: Handles error responses
        - `home.go`: Handles home page requests
    - `models/`
        - `Artist.go`: Defines the data structure for artists
    - `services/`
        - `APIClient.go`: Implements the API client for external data fetching
- `Templates/`
    - `detailsPage.html`: HTML template for the artist details page
    - `error.html`: HTML template for error pages
    - `index.html`: HTML template for the home page
- `static/`
    - `artist_location_cache.json`: Cached data for artist locations
- `go.mod`: Go module file
- `README.MD`: This file

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.16 or higher
- A modern web browser

## Installation

To install Groupe Tracker, follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/groupe-tracker.git
   ```
2. Navigate to the project directory:
   ```
   cd groupe-tracker
   ```
3. Install the dependencies:
   ```
   go mod tidy
   ```

## Usage

To run the Groupe Tracker application:

1. From the root directory of the project, run:
   ```
   go run cmd/server/main.go
   ```
2. Open your web browser and navigate to `http://localhost:8000` (or whatever port you've configured)

### Searching for Artists

1. Use the search bar at the top of the page to enter your search query
2. Select a search criteria from the dropdown menu (Name, Members, Creation Date, First Album, Location, or All)
3. As you type, autocomplete suggestions will appear
4. Press Enter or click the search button to perform the search

## Authors

* Faisal Almarzouqi - (falmarzo)
* Sayed Jalal - (sgallaf)
* Mudhi khalid - (mbaniham)
