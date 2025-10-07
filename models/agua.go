package models

import (
	"gorm.io/gorm"
)

type Agua struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Atividade string `json:"atividade"`
	Peso      int    `json:"peso"` // peso em kg
}

type PontosAgua struct {
	Agua      Agua
	Pontuacao int `json:"pontuacao"`
}

var agua []Agua

func QuantidadeAgua(a Agua) int {
	// 35ml por kg ou 0,001 por kg (dividir por 1000)
	totalPorDia := a.Peso * 35 / 1000
	println("Total de água por dia: ", totalPorDia, "L")
	return totalPorDia

}
