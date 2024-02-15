package main

import (
	"fmt"
	"mcserv2/internal/controller"
	"mcserv2/internal/router"
	"mcserv2/internal/service"
	"net/http"
)

func main() {
	sO := service.NewOrderService()
	cn := controller.NewHandleOrders(sO)
	r := router.StRout(cn)
	fmt.Println("server mcs2 start on :8082")
	http.ListenAndServe(":8082", r)
}
