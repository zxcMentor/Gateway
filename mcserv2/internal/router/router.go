package router

import (
	"github.com/go-chi/chi"
	"mcserv2/internal/controller"
)

func StRout(cn *controller.HandleOrders) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/orders/create", cn.CreateOrder)
	r.Get("/orders/get", cn.GetOrder)
	return r
}
