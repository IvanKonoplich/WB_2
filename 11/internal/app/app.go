package app

import (
	"WB2/11/internal/controller"
	"WB2/11/internal/infrastructure/storage"
	"WB2/11/internal/usecases"
)

// InitApp создает все необходимые объекты для работы сервера
func InitApp(port string) error {
	store := storage.NewStorage()
	uc := usecases.New(store)
	c := controller.New(uc)
	c.InitRouter()
	s := controller.Server{}
	return s.InitServer(port)
}
