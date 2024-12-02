package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "example.com/greetings/models"
	"github.com/gorilla/mux"
)

func GetAbgeordnete(w http.ResponseWriter, r *http.Request) {

	abgeordnete := models.GetAbgeordnete()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abgeordnete)
}

func GetAbgeordneteByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	abgeordneter := models.AbgeordneterById(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abgeordneter)
}

func CreateAbgeordnete(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var abgeordneter models.Abgeordnete_put
	err := decoder.Decode(&abgeordneter)
	if err != nil {
		panic(err)
	}

	models.PutAbgeordneter(abgeordneter)

	w.Header().Set("Content-Type", "application/json")
}

func GetAbgeordneteVote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	abgeordneter := models.GetAbgeordneteVote(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abgeordneter)
}
