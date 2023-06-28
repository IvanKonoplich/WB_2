package controller

import (
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("поступил новый запрос")
		f(w, r)
		log.Println("запрос обработан")
	}
}
