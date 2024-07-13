package repository

import (
	"github.com/google/uuid"
	"sync"
)

type ItemData struct {
	Name string `json:"name,omitempty"`
}

type Store struct {
	mu    sync.Mutex
	store map[uuid.UUID]ItemData
}

func NewStore() *Store {
	return &Store{
		store: make(map[uuid.UUID]ItemData),
	}
}

func (db *Store) SetData(id uuid.UUID, data ItemData) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.store[id] = data
}

func (db *Store) GetData(id uuid.UUID) ItemData {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, _ := db.store[id]

	return data
}

func (db *Store) DeleteData(id uuid.UUID) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.store, id)
}

func (db *Store) GetAllData() map[uuid.UUID]ItemData {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.store
}
