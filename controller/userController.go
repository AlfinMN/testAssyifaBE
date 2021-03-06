package controller

import (
	"assyifa/models"
	"assyifa/usecase"
	"assyifa/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func UserController(r *mux.Router, DataUser usecase.UserUsecase) {
	handler := UserHandler{DataUser}

	r.HandleFunc("/users", handler.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users", handler.AddUsers).Methods(http.MethodPost)
}

func (DataUser *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	data, err := DataUser.UserUsecase.GetAllData()

	if err != nil {
		w.Write([]byte("Data not found"))
	}
	var status models.Response
	status.Status = http.StatusOK
	status.Messages = "data Users Success"
	status.Data = data

	byteOfUser, err := json.Marshal(status)
	if err != nil {
		w.Write([]byte("Oops something wrong "))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfUser)
}

func (DataUser UserHandler) AddUsers(w http.ResponseWriter, r *http.Request) {

	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)

	var isTrue = utils.ValidationUser(&user)

	if err != nil {
		log.Println(err)
	}
	if isTrue == false {
		w.Write([]byte("Data Belum Lengkap "))
	} else {
		err = DataUser.UserUsecase.AddUser(user)
		if err != nil {
			log.Println(err)
		}
		log.Println("Insert successful")
		w.Write([]byte("Insert successful"))
	}
}
