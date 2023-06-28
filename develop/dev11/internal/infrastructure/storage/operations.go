package storage

import (
	"WB2/develop/11/internal/entities"
)

// CreateEvent добавляет событие в хранилище
func (s *Storage) CreateEvent(event entities.Event) error {
	s.hm[event.UserID] = append(s.hm[event.UserID], event)
	return nil
}

// CheckEvent проверяет есть ли в хранилище событие
func (s *Storage) CheckEvent(event entities.Event) (bool, error) {
	events, ok := s.hm[event.UserID]
	if !ok {
		return false, nil
	}
	for _, j := range events {
		if j == event {
			return true, nil
		}
	}
	return false, nil
}

// UpdateEvent меняет данные о событии в хранилище
func (s *Storage) UpdateEvent(event entities.Event) error {
	events := s.hm[event.UserID]
	for i, j := range events {
		if j.UserID == event.UserID && event.Date == j.Date {
			events[i] = event
		}
	}
	return nil
}

// DeleteEvent удаляет событие из хранилища
func (s *Storage) DeleteEvent(event entities.Event) error {
	events := s.hm[event.UserID]
	for i, j := range events {
		if j == event {
			remove(events, i)
		}
	}
	return nil
}

// EventsForDay возвращает список событий на указанный день
func (s *Storage) EventsForDay(event entities.Event) ([]entities.Event, error) {
	events := s.hm[event.UserID]
	result := make([]entities.Event, 0, len(events))
	for _, j := range events {
		if j.Date.Day() == event.Date.Day() {
			result = append(result, j)
		}
	}
	return result, nil
}

// EventsForWeek возвращает список событий на указанную неделю
func (s *Storage) EventsForWeek(event entities.Event) ([]entities.Event, error) {
	events := s.hm[event.UserID]
	result := make([]entities.Event, 0, len(events))
	for _, j := range events {
		jy, jw := j.Date.ISOWeek()
		ey, ew := event.Date.ISOWeek()
		if jy == ey && jw == ew {
			result = append(result, j)
		}
	}
	return result, nil
}

// EventsForMonth возвращает список событий на указанный месяц
func (s *Storage) EventsForMonth(event entities.Event) ([]entities.Event, error) {
	events := s.hm[event.UserID]
	result := make([]entities.Event, 0, len(events))
	for _, j := range events {
		if j.Date.Month() == event.Date.Month() {
			result = append(result, j)
		}
	}
	return result, nil
}

func remove(s []entities.Event, i int) []entities.Event {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
