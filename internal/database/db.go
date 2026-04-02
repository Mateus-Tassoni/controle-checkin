package database

import (
	"controle-checkin/internal/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBanco() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)

	for i := 1; i <= 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Tentativa %d: Banco ainda não respondeu Esperando 2s...", i)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Falha brutal ao conectar no banco de dados após 5 tentativas:", err)
	}

	err = DB.AutoMigrate(&models.Evento{}, &models.Convidado{})
	if err != nil {
		log.Fatal("Falha ao migrar tabelas:", err)
	}

	log.Println("PostgreSQL conectado e tabelas sincronizadas!")
}
