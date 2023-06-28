package controller

import "net/http"

// InitRouter инициализирует маршрутизацию запросов
func (c *Controller) InitRouter() {
	http.HandleFunc("/create_event", logging(c.handleCreateEvent))
	http.HandleFunc("/update_event", logging(c.handleUpdateEvent))
	http.HandleFunc("/delete_event", logging(c.handleDeleteEvent))
	http.HandleFunc("/events_for_day", logging(c.handleEventsForDay))
	http.HandleFunc("/events_for_week", logging(c.handleEventsForWeek))
	http.HandleFunc("/events_for_month", logging(c.handleEventsForMonth))
}
