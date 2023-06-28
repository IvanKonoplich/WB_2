package usecases

import (
	"WB2/develop/11/internal/entities"
)

type storage interface {
	CreateEvent(event entities.Event) error
	CheckEvent(event entities.Event) (bool, error)
	UpdateEvent(event entities.Event) error
	DeleteEvent(event entities.Event) error
	EventsForDay(event entities.Event) ([]entities.Event, error)
	EventsForWeek(event entities.Event) ([]entities.Event, error)
	EventsForMonth(event entities.Event) ([]entities.Event, error)
}

// UseCase структура реализующая методы usecase. Она зависит только от entities и через контроллер применяет методы для обработки запросов
type UseCase struct {
	storage storage
}

// New создает новый экземпляр UseCase
func New(storage storage) *UseCase {
	return &UseCase{
		storage: storage,
	}
}
