package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "example.com/greetings/models"
	"github.com/gorilla/mux"
)

func GetAbstimmungen(w http.ResponseWriter, r *http.Request) {

	abstimmungen := models.GetAbstimmungen()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abstimmungen)
}

func GetAbstimmungByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	abstimmung := models.AbstimmungById(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(abstimmung)
}

func CreateAbstimmung(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var abstimmung models.Abstimmung_put
	err := decoder.Decode(&abstimmung)
	if err != nil {
		panic(err)
	}

	models.PutAbstimmung(abstimmung)

	w.Header().Set("Content-Type", "application/json")
}

func CreateAbstimmungMulti(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var abstimmungen []models.Abstimmung_put
	err := decoder.Decode(&abstimmungen)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(abstimmungen); i++ {
		abstimmung := abstimmungen[i]
		models.PutAbstimmung(abstimmung)
	}

	w.Header().Set("Content-Type", "application/json")
}
