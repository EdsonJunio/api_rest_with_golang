package controllers

import (
	"api_rest_with_golang/database"
	"api_rest_with_golang/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornarUmaPersonalidades(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
