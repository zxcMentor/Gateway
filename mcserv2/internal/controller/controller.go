package controller

import (
	"encoding/json"
	"log"
	"mcserv2/internal/service"
	"net/http"
)

type HandleOrders struct {
	servOrder *service.OrderService
}

func NewHandleOrders(sO *service.OrderService) *HandleOrders {
	return &HandleOrders{sO}
}

func (h *HandleOrders) CreateOrder(w http.ResponseWriter, r *http.Request) {
	id := 1
	mess, err := h.servOrder.CreateOrder(id)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("order create")
	w.Write([]byte(mess))
}

func (h *HandleOrders) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := 1
	order, err := h.servOrder.GetOrder(id)
	if err != nil {
		log.Println("err", err)
	}
	jsData, err := json.Marshal(order)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("order get")
	w.Write(jsData)
}
