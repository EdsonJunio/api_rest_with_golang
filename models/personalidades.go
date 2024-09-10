package models

type Personalidade struct {
	Id        int    `json:"id"`
	Nome      string `json:"nome"`
	Historias string `json:"historias"`
}

var Personalidades []Personalidade
