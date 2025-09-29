package models

import (
	"gorm.io/gorm"
)

type Agua struct {
	gorm.Model
	Atividade string `json:"atividade"`
	Peso      int    `json:"peso"` // peso em kg
	Pontuacao int    `json:"pontuacao"`
}

var agua []Agua

func QuantidadeAgua(a Agua) int {
	// 35ml por kg ou 0,001 por kg (dividir por 1000)
	totalPorDia := a.Peso * 35 / 1000
	println("Total de água por dia: ", totalPorDia, "L")
	return totalPorDia

}

func ResetAgua() {
	var quantidadeDeAguaSemanal int
	if len(agua) > 0 {
		quantidadeDeAguaSemanal = QuantidadeAgua(agua[0]) * 7
		println("Total de água por semana: ", quantidadeDeAguaSemanal, "L")
	}
	for i := range agua {
		agua[i].Pontuacao = 0
		println("Pontuação de água resetada para: ", agua[i].Pontuacao)
	}
}
