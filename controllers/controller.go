package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariaeduardagsa-sys/go-gin/database"
	"github.com/mariaeduardagsa-sys/go-gin/models"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&trabalho)
	c.JSON(http.StatusOK, trabalho)
}

func GetTrabalhoPorId(c *gin.Context) {
	var trabalho models.Trabalho
	id := c.Params.ByName("id")
	database.DB.First(&trabalho, id)
	if trabalho.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Trabalho não encontrado"})
		return
	}
	c.JSON(http.StatusOK, trabalho)
}

func DeleteTrabalho(c *gin.Context) {
	var trabalho models.Trabalho
	id := c.Params.ByName("id")
	database.DB.Delete(&trabalho, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Trabalho deletado com sucesso"})
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

	database.DB.Model(&trabalho).UpdateColumns(trabalho)
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&academia)
	c.JSON(http.StatusOK, academia)
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

	database.DB.Model(&academia).UpdateColumns(academia)
	database.DB.Save(&academia)
	c.JSON(http.StatusOK, academia)
}
