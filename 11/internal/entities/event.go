package entities

import "time"

// Event - это основная сущность сервера. Содержит id отправителя, дату и текст с информацией о событии на это время
type Event struct {
	Date    time.Time
	UserID  int
	Message string
}
