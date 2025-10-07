package models

import (
	"gorm.io/gorm"
)

type Trabalho struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Atividade string `json:"atividade"`
	Status    string `json:"status"`
}

type PontosTrabalho struct {
	Trabalho  Trabalho
	Pontuacao int `json:"pontuacao"`
}

var Trabalhos []Trabalho
