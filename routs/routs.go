// src/routes/routes.go
package routes

import (
	controllers "example.com/greetings/controllers"
	"github.com/gorilla/mux" // One of the popular routing frameworks
)

// this is a comment
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")

	router.HandleFunc("/api/abgeordnete", controllers.GetAbgeordnete).Methods("GET")
	router.HandleFunc("/api/abgeordnete/{id}", controllers.GetAbgeordneteByID).Methods("GET")
	router.HandleFunc("/api/abgeordnete", controllers.CreateAbgeordnete).Methods("POST")
	router.HandleFunc("/api/abgeordnetevotes/{id}", controllers.GetAbgeordneteVote).Methods("GET")

	router.HandleFunc("/api/gesetze", controllers.GetGesetze).Methods("GET")
	router.HandleFunc("/api/gesetze/{id}", controllers.GetGesetzByID).Methods("GET")
	router.HandleFunc("/api/gesetze", controllers.CreateGesetz).Methods("POST")

	router.HandleFunc("/api/abstimmungen", controllers.GetAbstimmungen).Methods("GET")
	router.HandleFunc("/api/abstimmungen/{id}", controllers.GetAbstimmungByID).Methods("GET")
	router.HandleFunc("/api/abstimmungen", controllers.CreateAbstimmung).Methods("POST")
	//router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	//router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")

	return router
}
