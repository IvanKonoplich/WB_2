package app

import (
	controller2 "WB2/develop/dev11/internal/controller"
	"WB2/develop/dev11/internal/infrastructure/storage"
	"WB2/develop/dev11/internal/usecases"
)

// InitApp создает все необходимые объекты для работы сервера
func InitApp(port string) error {
	store := storage.NewStorage()
	uc := usecases.New(store)
	c := controller2.New(uc)
	c.InitRouter()
	s := controller2.Server{}
	return s.InitServer(port)
}
