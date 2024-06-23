package services

import (
	"encoding/json"
	"groupe-tracker/internal/models"
	"io/ioutil"
	"net/http"
	"strings"
)

type APIClient struct {
	baseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{baseURL: baseURL}
}

func (c *APIClient) GetAllArtists() ([]models.Artist, error) {
	resp, err := http.Get(c.baseURL + "/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var artists []models.Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func (c *APIClient) FetchRelationData(url string) (models.RelationData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.RelationData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.RelationData{}, err
	}

	var relationData models.RelationData
	if err := json.Unmarshal(body, &relationData); err != nil {
		return models.RelationData{}, err
	}

	return relationData, nil
}

func (c *APIClient) fetchData(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	if strings.Contains(url, "locations") {
		var data models.LocationData
		err = json.Unmarshal(body, &data)
		result = data
	} else if strings.Contains(url, "dates") {
		var data models.DatesData
		err = json.Unmarshal(body, &data)
		result = data
	} else if strings.Contains(url, "relation") {
		var data models.RelationData
		err = json.Unmarshal(body, &data)
		result = data
	} else {
		err = json.Unmarshal(body, &result)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}
