package controller

import (
	"WB2/develop/11/internal/entities"
)

type useCase interface {
	CreateEvent(event entities.Event) error
	UpdateEvent(event entities.Event) error
	DeleteEvent(event entities.Event) error
	EventsForDay(event entities.Event) ([]entities.Event, error)
	EventsForWeek(event entities.Event) ([]entities.Event, error)
	EventsForMonth(event entities.Event) ([]entities.Event, error)
}

// Controller - это абстракция контроллера содержащая usecase для вызова его методов обработчиками
type Controller struct {
	uc useCase
}

// New создает новый экземпляр контроллера
func New(uc useCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
