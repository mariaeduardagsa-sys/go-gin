package models

import (
	"gorm.io/gorm"
)

type Academia struct {
	gorm.Model
	Atividade  string `json:"atividade"`
	Status     string `json:"status"`
	Superiores bool   `json:"superiores"`
	Inferiores bool   `json:"inferiores"`
	Cardio     bool   `json:"cardio"`
	Pontuacao  int    `json:"pontuacao"`
}

var Exercicios []Academia
