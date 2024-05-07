package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL)) // hash the original URL
	data := hasher.Sum(nil)           // hash it
	fmt.Println(data)
	hash := hex.EncodeToString(data) // convert hash to string
	return hash[:8]                  // 8 characters
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL // for now
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}

	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func main() {
	fmt.Println("Hello, World!")
	OriginalURL := "https://www.google.com"
	fmt.Println(generateShortURL(OriginalURL))
}
