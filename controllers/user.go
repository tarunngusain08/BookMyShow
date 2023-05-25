package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dunzoit/BookMyShow/exceptions"
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/services"
	"github.com/gorilla/mux"
)

type User struct {
	service services.UserServices
}

func NewUserController(service services.UserServices) *User {
	return &User{
		service: service,
	}
}

func (u *User) GetUsers(w http.ResponseWriter, req *http.Request) {
	users, err := u.service.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (u *User) CreateUser(w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = u.service.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (u *User) GetUser(w http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(req)["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidUserID))
		return
	}

	user, err := u.service.GetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (u *User) UpdateUser(w http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(req)["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidUserID))
		return
	}

	updatedUser := &models.User{}
	err = json.NewDecoder(req.Body).Decode(updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = u.service.UpdateUser(userID, updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *User) DeleteUser(w http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(req)["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidUserID))
		return
	}

	err = u.service.DeleteUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
