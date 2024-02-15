package controller

import (
	"encoding/json"
	"log"
	"mcserv1/internal/service"
	"net/http"
)

type HandleUser struct {
	servUser *service.UserService
}

func NewHandleUser(su *service.UserService) *HandleUser {
	return &HandleUser{servUser: su}
}

func (h *HandleUser) CreateUser(w http.ResponseWriter, r *http.Request) {

	name := "diso"
	email := "@example"
	mess, err := h.servUser.CreateUser(name, email)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("user create")
	w.Write([]byte(mess))

}

func (h *HandleUser) GetUser(w http.ResponseWriter, r *http.Request) {
	name := "diso"
	user, err := h.servUser.GetUser(name)
	if err != nil {
		log.Println("err", err)
	}
	jsDt, err := json.Marshal(user)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("user get")
	w.Write(jsDt)
}
