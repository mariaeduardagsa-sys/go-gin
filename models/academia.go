package models

import (
	"gorm.io/gorm"
)

type Academia struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey"`
	Atividade  string `json:"atividade"`
	Status     string `json:"status"`
	Superiores bool   `json:"superiores"`
	Inferiores bool   `json:"inferiores"`
	Cardio     bool   `json:"cardio"`
}

type PontosAcademia struct {
	Academia  Academia
	Pontuacao int `json:"pontuacao"`
}

var Exercicios []Academia
