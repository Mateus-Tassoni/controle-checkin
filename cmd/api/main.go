package main

import (
	"controle-checkin/internal/database"
	"controle-checkin/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConectarBanco()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/checkin", handlers.RealizarCheckIn)
		api.POST("/convidados", handlers.CriarConvidado)
		api.GET("/convidados", handlers.ListarConvidados)
	}

	log.Println("Servidor rodando na porta 8080...")
	r.Run(":8080")
}
