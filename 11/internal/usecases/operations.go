package usecases

import (
	"WB2/11/internal/entities"
	"errors"
)

// CreateEvent проверяет существование события и если его нет, то создает его
func (uc *UseCase) CreateEvent(event entities.Event) error {
	ok, err := uc.storage.CheckEvent(event)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("событие уже существует")
	}
	return uc.storage.CreateEvent(event)
}

// UpdateEvent проверяет существование события и если оно есть, то меняет его
func (uc *UseCase) UpdateEvent(event entities.Event) error {
	ok, err := uc.storage.CheckEvent(event)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("событие не создано")
	}
	return uc.storage.UpdateEvent(event)
}

// DeleteEvent проверяет существование события и если оно есть, то удаляет его
func (uc *UseCase) DeleteEvent(event entities.Event) error {
	ok, err := uc.storage.CheckEvent(event)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("событие не создано")
	}
	return uc.storage.DeleteEvent(event)
}

// EventsForDay возвращает список событий на день
func (uc *UseCase) EventsForDay(event entities.Event) ([]entities.Event, error) {
	return uc.storage.EventsForDay(event)
}

// EventsForWeek возвращает список событий на неделю
func (uc *UseCase) EventsForWeek(event entities.Event) ([]entities.Event, error) {
	return uc.storage.EventsForWeek(event)
}

// EventsForMonth возвращает список событий на месяц
func (uc *UseCase) EventsForMonth(event entities.Event) ([]entities.Event, error) {
	return uc.storage.EventsForMonth(event)
}
