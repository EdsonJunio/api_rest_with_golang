package routes

import (
	"api_rest_with_golang/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaNovaPersonalidades).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
