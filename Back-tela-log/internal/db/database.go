package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	err := godotenv.Load("C:/Users/heito/OneDrive/Documentos/Programas/Go-back-projects/Back-tela-log/config/login.env")

	if err != nil {
		log.Fatal("Erro ao carregar .env", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, password, name)

	db, err := sql.Open("mysql", dns)

	if err != nil {
		log.Fatal("Erro ao abrir db!", err)
	}

	errs := db.Ping()

	if errs != nil {
		log.Fatal("Erro ao conectar ao banco!", errs)
	}

	return db
}
