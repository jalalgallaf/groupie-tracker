package services

import (
	"encoding/json"
	"fmt"
	"groupe-tracker/internal/models"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CacheItem struct {
	Data      interface{}
	ExpiresAt time.Time
}

var cacheMutex sync.RWMutex

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

	//get locations for each artist
	for i, artist := range artists {
		locationData, err := c.GetArtestLocation(&artist)
		if err != nil {
			return nil, err
		}
		artists[i].LocationData = locationData
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

func (c *APIClient) GetArtestLocation(artist *models.Artist) (interface{}, error) {

	cacheFile := "artist_location_cache.json"
	cacheExpiration := 10 * time.Minute

	cacheData, err := loadFromCache(cacheFile, artist.ID)
	if err == nil {
		return cacheData, nil
	}

	locationAPILink := artist.Locations
	resp, err := http.Get(locationAPILink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data models.LocationData
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	saveToCache(cacheFile, artist.ID, data.Locations, cacheExpiration)

	return data.Locations, nil

}

func loadFromCache(file string, artistId int) (interface{}, error) {

	cacheMutex.RLock()
	defer cacheMutex.RUnlock()

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cache map[string]CacheItem
	err = json.Unmarshal(data, &cache)

	if err != nil {
		return nil, err
	}

	item, found := cache[strconv.Itoa(artistId)]

	if !found || time.Now().After(item.ExpiresAt) {
		return nil, fmt.Errorf("Cache miss of expired")
	}

	return item.Data, nil
}

func saveToCache(file string, artistId int, data interface{}, expiration time.Duration) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	var cache map[string]CacheItem

	fileData, err := os.ReadFile(file)

	if err == nil {
		json.Unmarshal(fileData, &cache)
	}
	if cache == nil {
		cache = make(map[string]CacheItem)
	}

	cache[strconv.Itoa(artistId)] = CacheItem{
		Data:      data,
		ExpiresAt: time.Now().Add(expiration),
	}

	newData, err := json.Marshal(cache)

	if err != nil {
		fmt.Printf("Error marshaling cache: %v\n", err)
		return

	}

	err = os.WriteFile(file, newData, 0644)
	if err != nil {
		fmt.Printf("Error writing cache file: %v\n", err)
	}

}
