package router

import (
	"geteway/internal/middleware"
	"github.com/go-chi/chi"
	"net/http"
)

func StRout() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Reverse)
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {})
	return r
}
