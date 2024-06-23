# Groupie Tracker

Groupie Tracker is a web application that displays information about various music artists and bands. It uses data from the Groupie Tracker API to show details such as band members, creation date, first album, and more.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [API Integration](#api-integration)
- [Contributing](#contributing)
- [License](#license)

## Features

- Display a grid of artists/bands with their images
- Search functionality to filter artists by name, members, creation date, or first album
- Detailed view for each artist/band
- Responsive design for various screen sizes
- Error handling with custom error pages

## Technologies Used

- Go (Golang) for backend logic
- HTML and CSS for frontend
- Groupie Tracker API for data

## Installation

1. Ensure you have Go installed on your system. If not, download and install it from [golang.org](https://golang.org/).

2. Clone this repository:
   ```
   git clone https://github.com/yourusername/groupie-tracker.git
   ```

3. Navigate to the project directory:
   ```
   cd groupie-tracker
   ```


## Usage

1. Run the application:
   ```
   go run cmd/server/main.go
   ```

2. Open a web browser and navigate to `http://localhost:3000`

3. Use the search bar to filter artists by name, members, creation date, or first album.

4. Click on an artist's card to view more details about them.

## Project Structure

```
groupie-tracker/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   │   ├── artist.go
│   │   ├── error.go
│   │   └── home.go
│   │
│   ├── models/
│   │   └── artist.go
│   │
│   ├── services/
│   │   └── api_client.go
│   │
│   └── templates/
│       ├── index.html
│       ├── detailsPage.html
│       └── error.html
│
├── static/
│   └── css/
│       └── style.css
│
├── .gitignore
├── go.mod
└── README.md
```