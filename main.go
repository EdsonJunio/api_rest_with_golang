package main

import (
	"api_rest_with_golang/models"
	"api_rest_with_golang/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome1", Historias: "historia 2"},
		{Id: 2, Nome: "nome3", Historias: "historia 4"},
	}

	routes.HandleResquest()
}
