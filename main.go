package main

import (
	"api_rest_with_golang/models"
	"api_rest_with_golang/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "nome1", Historias: "historia 2"},
		{Nome: "nome3", Historias: "historia 4"},
	}

	fmt.Println("teste")
	routes.HandleResquest()
}
