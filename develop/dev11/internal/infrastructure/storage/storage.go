package storage

import (
	"WB2/develop/11/internal/entities"
)

// Storage - это абстракция хранилища, которая в данном случае реализуется через map
type Storage struct {
	hm map[int][]entities.Event
}

// NewStorage создает новый экземпляр хранилища
func NewStorage() *Storage {
	return &Storage{hm: make(map[int][]entities.Event, 100)}
}
