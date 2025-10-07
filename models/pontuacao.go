package models

import (
	"gorm.io/gorm"
)

type Pontuacao struct {
	gorm.Model
	ID        uint `json:"id" gorm:"primaryKey"`
	Pontuacao int  `json:"pontuacao"`
}

func IncrementaPontuacaoTrabalho(p *PontosTrabalho) {
	p.Pontuacao += 10
	println("Pontuação atualizada para: ", p.Pontuacao)
}

func DecrementaPontuacaoTrabalho(p *PontosTrabalho) {
	p.Pontuacao -= 10
	println("Pontuação atualizada para: ", p.Pontuacao)
}

func IncrementaPontuacaoAcademia(p *PontosAcademia) {
	p.Pontuacao += 20
	println("Pontuação da academia atualizada para:", p.Pontuacao)
}

func DecrementaPontuacaoAcademia(p *PontosAcademia) {
	p.Pontuacao -= 20
	println("Pontuação atualizada para: ", p.Pontuacao)
}

func IncrementaPontuacaoAgua(p *PontosAgua) {
	p.Pontuacao += 1
	println("Pontuação da água atualizada para:", p.Pontuacao)
}

func DecrementaPontuacaoAgua(p *PontosAgua) {
	if p.Pontuacao > 0 {
		p.Pontuacao -= 1
	}
	println("Pontuação da água atualizada para:", p.Pontuacao)
}

func ResetAgua(pontosAgua []PontosAgua) {
	if len(pontosAgua) == 0 {
		println("Nenhum registro de água encontrado.")
		return
	}

	quantidadeDeAguaSemanal := QuantidadeAgua(pontosAgua[0].Agua) * 7

	for i := range pontosAgua {
		if pontosAgua[i].Pontuacao >= quantidadeDeAguaSemanal {
			pontosAgua[i].Pontuacao = 0
			println("Pontuação resetada! Agora:", pontosAgua[i].Pontuacao)
		} else {
			println("Pontuação atual:", pontosAgua[i].Pontuacao)
		}
	}
}
