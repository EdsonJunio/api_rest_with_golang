package main

import (
	"api_rest_with_golang/database"
	"api_rest_with_golang/routes"
	"fmt"
)

func main() {

	database.ConectarComObancoDedados()

	fmt.Println("teste")
	routes.HandleResquest()
}
