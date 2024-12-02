package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "example.com/greetings/models"
	"github.com/gorilla/mux"
)

func GetGesetze(w http.ResponseWriter, r *http.Request) {

	gesetze := models.GetGesetze()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gesetze)
}

func GetGesetzByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	gesetz := models.GesetzById(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gesetz)
}

func CreateGesetz(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var gesetz models.Gesetz_put
	err := decoder.Decode(&gesetz)
	if err != nil {
		panic(err)
	}

	models.PutGesetz(gesetz)

	w.Header().Set("Content-Type", "application/json")
}
