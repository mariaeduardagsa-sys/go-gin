package main

import (
	"github.com/mariaeduardagsa-sys/go-gin/database"
	"github.com/mariaeduardagsa-sys/go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	// models.Trabalhos = []models.Trabalho{
	// 	{Atividade: "Trabalho", Status: "success"},
	// 	{Atividade: "Trabalho", Status: "success"},
	// 	{Atividade: "Trabalho", Status: "pending"},
	// }
	routes.HandleRequests()
}
