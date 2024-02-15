package main

import (
	"fmt"
	"mcserv1/internal/controller"
	"mcserv1/internal/router"
	"mcserv1/internal/service"
	"net/http"
)

func main() {
	su := service.NewUserService()
	cn := controller.NewHandleUser(su)
	r := router.StRout(cn)
	fmt.Println("server mcs1 start on :8081")
	http.ListenAndServe(":8081", r)
}
