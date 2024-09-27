package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectarComObancoDedados() {
	stringConexao := "host=localhost port=5432 user=root dbname=root password=root sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("erro ao conectar com o banco de dados")
	}
}
