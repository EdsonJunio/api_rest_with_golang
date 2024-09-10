package controllers

import (
	"api_rest_with_golang/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
