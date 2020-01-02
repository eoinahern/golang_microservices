package services

import (
	"net/http"

	"github.com/eoinahern/golang_microservices/microservice_project/domain"
	"github.com/eoinahern/golang_microservices/microservice_project/utils"
)

type itemService struct {
}

var ItemService itemService = itemService{}

func (i *itemService) GetItem(itemID int) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement please",
		StatusCode: http.StatusNotFound,
	}
}
