package handlers

import (
	"controle-checkin/internal/database"
	"controle-checkin/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type CheckInRequest struct {
	CodigoQR string `json:"codigo_qr" binding:"required"`
}

func RealizarCheckIn(c *gin.Context) {
	var req CheckInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "QR Code inválido"})
		return
	}

	tx := database.DB.Begin()
	var convidado models.Convidado

	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("codigo_qr = ?", req.CodigoQR).
		First(&convidado).Error

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"erro": "Ingresso não encontrado"})
		return
	}

	if convidado.Status == "CHECKED_IN" {
		tx.Rollback()
		c.JSON(http.StatusConflict, gin.H{
			"erro": "Acesso Negado: Ingresso já utilizado!",
			"nome": convidado.Nome,
		})
		return
	}

	convidado.Status = "CHECKED_IN"

	if err := tx.Save(&convidado).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro no banco"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Acesso Liberado",
		"nome":     convidado.Nome,
	})
}
