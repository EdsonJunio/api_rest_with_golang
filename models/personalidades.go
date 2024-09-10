package models

type Personalidade struct {
	Nome      string `json:"nome"`
	Historias string `json:"historias"`
}

var Personalidades []Personalidade
