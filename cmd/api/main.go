package main

import (
	"controle-checkin/internal/database"
	"controle-checkin/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Conecta no Postgres
	database.ConectarBanco()

	// 2. Prepara a API
	r := gin.Default()

	// 3. Cria as rotas que a gente precisa
	api := r.Group("/api")
	{
		api.POST("/checkin", handlers.RealizarCheckIn)

		// OLHA AS ROTAS AQUI QUE FALTARAM NO SEU CÓDIGO:
		api.POST("/convidados", handlers.CriarConvidado)
		api.GET("/convidados", handlers.ListarConvidados)
	}

	// 4. Inicia o servidor
	log.Println("Servidor rodando na porta 8080...")
	r.Run(":8080")
}
