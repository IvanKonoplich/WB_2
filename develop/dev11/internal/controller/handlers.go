package controller

import (
	"WB2/develop/11/internal/entities"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (c *Controller) handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, 400, nil, "incorrect method")
		return
	}
	event := makeEventFromQuery(w, r)
	if err := c.uc.CreateEvent(event); err != nil {
		errorResponse(w, 503, err, "error while creating event: ")
		return
	}
}

func (c *Controller) handleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, 400, nil, "incorrect method")
		return
	}
	event := makeEventFromQuery(w, r)
	if err := c.uc.UpdateEvent(event); err != nil {
		errorResponse(w, 503, err, "error while updating event: ")
		return
	}
}

func (c *Controller) handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, 400, nil, "incorrect method")
		return
	}
	event := makeEventFromQuery(w, r)
	if err := c.uc.DeleteEvent(event); err != nil {
		errorResponse(w, 503, err, "error while deleting event: ")
		return
	}
}
func (c *Controller) handleEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, 400, nil, "incorrect method")
		return
	}
	event := makeEventFromQuery(w, r)
	result, err := c.uc.EventsForDay(event)
	if err != nil {
		errorResponse(w, 400, err, "error while getting events")
		return
	}
	sendResult(w, result)

}
func (c *Controller) handleEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, 400, nil, "incorrect method")
		return
	}
	event := makeEventFromQuery(w, r)
	result, err := c.uc.EventsForWeek(event)
	if err != nil {
		errorResponse(w, 400, err, "error while getting events")
		return
	}
	sendResult(w, result)
}
func (c *Controller) handleEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, 400, nil, "incorrect method")
		return
	}
	event := makeEventFromQuery(w, r)
	result, err := c.uc.EventsForMonth(event)
	if err != nil {
		errorResponse(w, 400, err, "error while getting events")
		return
	}
	sendResult(w, result)
}

func errorResponse(w http.ResponseWriter, sc int, err error, message string) {
	log.Println(message + err.Error())
	data, err := json.Marshal(responseError{Error: message + err.Error()})
	if err != nil {
		log.Printf("ошибка во время отправки ответа: %s\n", err.Error())
		return
	}
	w.WriteHeader(sc)
	if _, err = w.Write(data); err != nil {
		log.Printf("ошибка во время отправки ответа: %s\n", err.Error())
		return
	}
	log.Println("отправлен ответ с ошибкой")
	return
}

func makeEventFromQuery(w http.ResponseWriter, r *http.Request) entities.Event {
	switch r.Method {
	case http.MethodPost:
		var inc incomingQuery
		data := make([]byte, 100)
		n, err := r.Body.Read(data)
		if err != nil && err != io.EOF {
			errorResponse(w, 503, err, "error while reading body: ")
			return entities.Event{}
		}
		err = json.Unmarshal(data[:n], &inc)
		if err != nil {
			errorResponse(w, 503, err, "error while reading body: ")
			return entities.Event{}
		}
		if inc.UserID == 0 {
			errorResponse(w, 400, nil, "incorrect user_id")
			return entities.Event{}
		}
		var event entities.Event
		event.UserID = inc.UserID
		event.Message = inc.Message
		event.Date, err = time.Parse("2006-01-02", inc.Date)
		if err != nil {
			errorResponse(w, 400, err, "incorrect date format")
			return entities.Event{}
		}
		log.Printf("новый event %s\n", fmt.Sprint(event))
		return event
	case http.MethodGet:
		event := entities.Event{}
		userID, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			errorResponse(w, 503, err, "incorrect user id: ")
			return entities.Event{}
		}
		event.UserID = userID
		date, err := time.Parse("2006-01-02", r.FormValue("date"))
		if err != nil {
			errorResponse(w, 400, err, "incorrect date format")
			return entities.Event{}
		}
		event.Date = date
		return event
	}
	return entities.Event{}
}

func sendResult(w http.ResponseWriter, events []entities.Event) {
	var responseMessage string
	for _, j := range events {
		responseMessage += fmt.Sprint(j.Date) + ": " + j.Message + "\n"
	}
	data, err := json.Marshal(response{Response: responseMessage})
	if err != nil {
		log.Printf("ошибка во время отправки ответа: %s\n", err.Error())
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Printf("ошибка во время отправки ответа: %s\n", err.Error())
		return
	}
}
