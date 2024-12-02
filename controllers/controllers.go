package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "example.com/greetings/models"
	"github.com/gorilla/mux"
)

//var users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := models.Users()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	user := models.UserById(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user models.User_put
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	models.PutUser(user)

	w.Header().Set("Content-Type", "application/json")
}
