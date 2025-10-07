package database

import (
	"log"

	"github.com/mariaeduardagsa-sys/go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Trabalho{})
	DB.AutoMigrate(&models.Academia{})
	DB.AutoMigrate(&models.Agua{})
	DB.AutoMigrate(&models.Pontuacao{})
}
