package models

import (
	"gorm.io/gorm"
)

type Trabalho struct {
	gorm.Model
	Atividade string `json:"atividade"`
	Status    string `json:"status"`
	Pontuacao int    `json:"pontuacao"`
}

var Trabalhos []Trabalho
