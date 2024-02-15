package service

import (
	"fmt"
	"mcserv2/internal/models"
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) CreateOrder(id int) (string, error) {
	return fmt.Sprintf("create order: %d", id), nil
}

func (o *OrderService) GetOrder(id int) (models.Order, error) {
	return models.Order{
		ID:       id,
		UserName: "Gary",
		Items: []models.Item{
			{
				ProductName: "Cola",
				Quantity:    1,
			},
		},
	}, nil
}
