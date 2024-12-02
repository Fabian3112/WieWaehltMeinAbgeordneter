package main

import (
	"net/http"

	models "example.com/greetings/models"
	routes "example.com/greetings/routs"
)

func main() {
	models.Init_db()
	var router = routes.SetupRoutes()
	http.ListenAndServe(":8080", router)
}
