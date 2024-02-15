package router

import (
	"github.com/go-chi/chi"
	"mcserv1/internal/controller"
)

func StRout(cn *controller.HandleUser) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/users/create", cn.CreateUser)
	r.Get("/users/get", cn.GetUser)
	return r
}
