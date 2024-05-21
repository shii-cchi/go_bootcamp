package db

import (
	"github.com/google/uuid"
	"net/http"
	"sync"
)

type JsonData struct {
	Name string `json:"name"`
}

type Database struct {
	mu    sync.Mutex
	store map[uuid.UUID]JsonData
}

func NewDatabase() *Database {
	return &Database{
		store: make(map[uuid.UUID]JsonData),
	}
}

func (db *Database) setData(id uuid.UUID, data JsonData) int {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.store[id]

	if exists {
		db.store[id] = data
		return http.StatusOK
	}

	db.store[id] = data
	return http.StatusCreated
}

func (db *Database) getData(id uuid.UUID) (int, JsonData) {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, exists := db.store[id]

	if !exists {
		return http.StatusNotFound, JsonData{}
	}

	return http.StatusOK, data
}

func (db *Database) deleteData(id uuid.UUID) (int, string) {

}
