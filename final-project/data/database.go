package data

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var (
	store      *InMemoryStore
	storeMutex 	sync.Mutex
)

type InMemoryStore struct {
	Books     map[int]Book     `json:"books"`
	Authors   map[int]Author   `json:"authors"`
	Customers map[int]Customer `json:"customers"`
	Orders    map[int]Order    `json:"orders"`
}

func NewInMemoryStore() *InMemoryStore {
	if store == nil {
		store = &InMemoryStore{
			Books:     make(map[int]Book),
			Authors:   make(map[int]Author),
			Customers: make(map[int]Customer),
			Orders:    make(map[int]Order),
		}
		if err := loadFromJSONFile("database.json"); err != nil {
			fmt.Println("Error loading JSON:", err)
		}
	}
	return store
}

func loadFromJSONFile(filename string) error {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(store)
}

func saveToJSONFile(filename string, source *InMemoryStore) error {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(source)
}
