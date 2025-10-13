package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariaeduardagsa-sys/go-gin/database"
	"github.com/mariaeduardagsa-sys/go-gin/models"
	"gorm.io/gorm"
)

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E aí " + nome + ", tudo bem?",
	})
}

func GetTrabalho(c *gin.Context) {
	var todosTrabalhos []models.Trabalho
	database.DB.Find(&todosTrabalhos)
	c.JSON(http.StatusOK, todosTrabalhos)
}

func CreateTrabalho(c *gin.Context) {
	var trabalho models.Trabalho
	if err := c.ShouldBindJSON(&trabalho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&trabalho).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var pontuacaoTabela models.Pontuacao
	if err := database.DB.Where("atividade = ?", "Trabalho").First(&pontuacaoTabela).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			pontuacaoTabela = models.Pontuacao{
				Atividade: "Trabalho",
				Pontuacao: 0,
			}
			database.DB.Create(&pontuacaoTabela)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	database.DB.First(&pontuacaoTabela, "atividade = ?", trabalho.Atividade)
	database.DB.Create(&pontuacaoTabela)
	pontuacaoTabela.Atividade = "Trabalho"
	pontuacaoTabela.Pontuacao += 10
	database.DB.Save(&pontuacaoTabela)

	pontos := models.PontosTrabalho{Trabalho: trabalho, Pontuacao: pontuacaoTabela.Pontuacao}
	models.IncrementaPontuacaoTrabalho(&pontos)

	c.JSON(http.StatusOK, gin.H{
		"mensagem":  "Trabalho criado com pontuação atualizada",
		"trabalho":  trabalho,
		"pontuacao": pontuacaoTabela.Pontuacao,
	})
}

func GetTrabalhoPorId(c *gin.Context) {
	var trabalho models.Trabalho
	id := c.Params.ByName("id")
	database.DB.First(&trabalho, id)

	if trabalho.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trabalho não encontrado"})
		return
	}

	c.JSON(http.StatusOK, trabalho)
}

func DeleteTrabalho(c *gin.Context) {
	var trabalho models.Trabalho
	id := c.Params.ByName("id")
	database.DB.Delete(&trabalho, id)
	database.DB.Unscoped().Where("id = ?", id).Delete(&trabalho)

	c.JSON(http.StatusOK, gin.H{
		"message": "Trabalho deletado com sucesso"})

	var pontuacaoTabela models.Pontuacao
	database.DB.First(&pontuacaoTabela, "atividade = ?", trabalho.Atividade)
	database.DB.Create(&pontuacaoTabela)
	pontuacaoTabela.Atividade = "Trabalho"
	pontuacaoTabela.Pontuacao -= 10
	database.DB.Save(&pontuacaoTabela)

	pontos := models.PontosTrabalho{Trabalho: trabalho, Pontuacao: 0}
	models.DecrementaPontuacaoTrabalho(&pontos)
}

func EditaTrabalho(c *gin.Context) {
	var trabalho models.Trabalho
	id := c.Params.ByName("id")
	database.DB.First(&trabalho, id)
	if err := c.ShouldBindJSON(&trabalho); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&trabalho).Updates(trabalho)
	database.DB.Save(&trabalho)
	c.JSON(http.StatusOK, trabalho)
}

func GetAcademia(c *gin.Context) {
	var todosExercicios []models.Academia
	database.DB.Find(&todosExercicios)
	c.JSON(http.StatusOK, todosExercicios)
}

func CreateExercicio(c *gin.Context) {
	var academia models.Academia
	if err := c.ShouldBindJSON(&academia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&academia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var pontuacaoTabela models.Pontuacao
	if err := database.DB.Where("atividade = ?", "Academia").First(&pontuacaoTabela).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			pontuacaoTabela = models.Pontuacao{
				Atividade: "Academia",
				Pontuacao: 0,
			}
			database.DB.Create(&pontuacaoTabela)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	database.DB.First(&pontuacaoTabela, "atividade = ?", academia.Atividade)
	database.DB.Create(&pontuacaoTabela)
	pontuacaoTabela.Atividade = "Academia"
	pontuacaoTabela.Pontuacao += 20
	database.DB.Save(&pontuacaoTabela)

	pontos := models.PontosAcademia{Academia: academia, Pontuacao: pontuacaoTabela.Pontuacao}
	models.IncrementaPontuacaoAcademia(&pontos)

	c.JSON(http.StatusOK, gin.H{
		"mensagem":  "Exercício criado com pontuação atualizada",
		"academia":  academia,
		"pontuacao": pontuacaoTabela.Pontuacao,
	})
}

func GetExercicioPorId(c *gin.Context) {
	var academia models.Academia
	id := c.Params.ByName("id")
	database.DB.First(&academia, id)
	if academia.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Exercício não encontrado"})
		return
	}
	c.JSON(http.StatusOK, academia)
}

func DeleteExercicio(c *gin.Context) {
	var academia models.Academia
	id := c.Params.ByName("id")
	database.DB.Delete(&academia, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Exercício deletado com sucesso"})

	pontos := models.PontosAcademia{Academia: academia, Pontuacao: 0}
	models.DecrementaPontuacaoAcademia(&pontos)

	var pontuacaoTabela models.Pontuacao
	database.DB.First(&pontuacaoTabela, "atividade = ?", academia.Atividade)
	database.DB.Create(&pontuacaoTabela)
	pontuacaoTabela.Atividade = "Academia"
	pontuacaoTabela.Pontuacao -= 20
	database.DB.Save(&pontuacaoTabela)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Exercício deletado com sucesso",
		"pontuacao": pontos.Pontuacao,
	})
}

func EditaAcademia(c *gin.Context) {
	var academia models.Academia
	id := c.Params.ByName("id")
	database.DB.First(&academia, id)
	if err := c.ShouldBindJSON(&academia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&academia).Updates(academia)
	database.DB.Save(&academia)
	c.JSON(http.StatusOK, academia)
}

func GetAgua(c *gin.Context) {
	var agua []models.Agua
	result := database.DB.Find(&agua)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, agua)
}

func CreateAgua(c *gin.Context) {
	var agua models.Agua
	if err := c.ShouldBindJSON(&agua); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&agua).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	quantidade := models.QuantidadeAgua(agua)

	var pontuacaoTabela models.Pontuacao
	if err := database.DB.Where("atividade = ?", "Agua").First(&pontuacaoTabela).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			pontuacaoTabela = models.Pontuacao{
				Atividade: "Agua",
				Pontuacao: 0,
			}
			database.DB.Create(&pontuacaoTabela)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	database.DB.First(&pontuacaoTabela, "atividade = ?", agua.Atividade)
	database.DB.Create(&pontuacaoTabela)
	pontuacaoTabela.Atividade = "Agua"
	pontuacaoTabela.Pontuacao += 2
	database.DB.Save(&pontuacaoTabela)

	pontos := models.PontosAgua{Agua: agua, Pontuacao: pontuacaoTabela.Pontuacao}
	quantidadeSemanal := models.QuantidadeAgua(agua) * 7

	if pontos.Pontuacao >= quantidadeSemanal {
		models.ResetAgua([]models.PontosAgua{pontos})
		pontuacaoTabela.Pontuacao = 0
		database.DB.Save(&pontuacaoTabela)
		c.JSON(http.StatusOK, gin.H{
			"mensagem":  "Pontuação semanal atingida! Pontuação resetada.",
			"agua":      agua,
			"pontuacao": 0,
		})
		return
	}

	msg := fmt.Sprintf("Quantidade de água ideal por dia: %d L", quantidade)

	c.JSON(http.StatusOK, gin.H{
		"mensagem":  msg,
		"agua":      agua,
		"pontuacao": pontuacaoTabela.Pontuacao,
	})
}

func DeleteAguaById(c *gin.Context) {
	var agua models.Agua
	id := c.Params.ByName("id")
	database.DB.First(&agua, id)
	if agua.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Registro de água não encontrado"})
		return
	}
	database.DB.Delete(&agua, id)
	database.DB.Unscoped().Where("id = ?", id).Delete(&agua)

	var pontuacaoTabela models.Pontuacao
	database.DB.First(&pontuacaoTabela, "atividade = ?", agua.Atividade)
	database.DB.Create(&pontuacaoTabela)
	pontuacaoTabela.Atividade = "Agua"
	pontuacaoTabela.Pontuacao -= 2
	database.DB.Save(&pontuacaoTabela)

	if agua.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Registro de água não encontrado"})
		return
	}
	msg := "Registro de água deletado com sucesso"

	pontos := models.PontosAgua{Agua: agua, Pontuacao: 0}
	models.DecrementaPontuacaoAgua(&pontos)
	c.JSON(http.StatusOK, gin.H{
		"message":   msg,
		"pontuacao": pontos.Pontuacao,
	})
}

func EditaAgua(c *gin.Context) {
	var agua models.Agua
	id := c.Params.ByName("id")
	database.DB.First(&agua, id)
	if err := c.ShouldBindJSON(&agua); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&agua).Updates(agua)
	database.DB.Save(&agua)
	c.JSON(http.StatusOK, agua)
}
