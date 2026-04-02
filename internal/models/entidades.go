package models

import "gorm.io/gorm"

type Evento struct {
	gorm.Model
	Nome       string      `json:"nome"`
	Data       string      `json:"data"`
	Capacidade int         `json:"capacidade"`
	Convidados []Convidado `json:"convidados,omitempty" gorm:"foreignKey:EventoID"`
}

type Convidado struct {
	gorm.Model
	Nome     string `json:"nome"`
	CPF      string `json:"cpf" gorm:"uniqueIndex"`
	EventoID uint   `json:"evento_id"`
	CodigoQR string `json:"codigo_qr" gorm:"uniqueIndex"`
	Status   string `json:"status" gorm:"default:'PENDENTE'"`
}
