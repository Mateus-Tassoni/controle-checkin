package database

import (
	"controle-checkin/internal/models"
	"fmt"
	"log"
	"os"
	"time" // PRECISA DESSE CARA AQUI

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBanco() {
	// Carrega as variáveis do arquivo .env (se ele existir)
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
	}

	// Puxa as informações de forma segura
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Monta a string de conexão
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)

	// --- LOGICA DE RETRY (TENTAR NOVAMENTE) ---
	// Tentamos 5 vezes com um intervalo de 2 segundos entre elas
	for i := 1; i <= 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break // Se conectou, sai do loop e segue a vida
		}

		log.Printf("Tentativa %d: Banco ainda não respondeu (Postgres acordando...). Esperando 2s...", i)
		time.Sleep(2 * time.Second)
	}

	// Se depois de 5 vezes ainda der erro, aí sim a gente desiste
	if err != nil {
		log.Fatal("Falha brutal ao conectar no banco de dados após 5 tentativas:", err)
	}

	// Migra as tabelas
	err = DB.AutoMigrate(&models.Evento{}, &models.Convidado{})
	if err != nil {
		log.Fatal("Falha ao migrar tabelas:", err)
	}

	log.Println("PostgreSQL conectado e tabelas sincronizadas!")
}
