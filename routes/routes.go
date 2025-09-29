package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mariaeduardagsa-sys/go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:nome", controllers.Saudacao)

	r.GET("/trabalho", controllers.GetTrabalho)
	r.POST("/trabalho", controllers.CreateTrabalho)
	r.GET("/trabalho/:id", controllers.GetTrabalhoPorId)
	r.DELETE("/trabalho/:id", controllers.DeleteTrabalho)
	r.PATCH("/trabalho/:id", controllers.EditaTrabalho)

	r.GET("/academia", controllers.GetAcademia)
	r.POST("/academia", controllers.CreateExercicio)
	r.GET("/academia/:id", controllers.GetExercicioPorId)
	r.DELETE("/academia/:id", controllers.DeleteExercicio)
	r.PATCH("/academia/:id", controllers.EditaAcademia)

	r.GET("/agua", controllers.GetAgua)
	r.POST("/agua", controllers.CreateAgua)
	r.DELETE("/agua/reset", controllers.DeleteAgua)

	r.Run()
}
