package db

import (
	"github.com/google/uuid"
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

func (db *Database) SetData(id uuid.UUID, data JsonData) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.store[id] = data
}

func (db *Database) GetData(id uuid.UUID) JsonData {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, _ := db.store[id]

	return data
}

func (db *Database) DeleteData(id uuid.UUID) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.store, id)
}
