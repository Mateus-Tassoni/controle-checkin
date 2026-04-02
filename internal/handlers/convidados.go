package handlers

import (
	"controle-checkin/internal/database"
	"controle-checkin/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriarConvidado(c *gin.Context) {
	var convidado models.Convidado

	if err := c.ShouldBindJSON(&convidado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	if err := database.DB.Create(&convidado).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao salvar. Verifique se o CPF ou QR Code já estão cadastrados."})
		return
	}

	c.JSON(http.StatusCreated, convidado)
}

func ListarConvidados(c *gin.Context) {
	var convidados []models.Convidado

	database.DB.Find(&convidados)

	c.JSON(http.StatusOK, convidados)
}
