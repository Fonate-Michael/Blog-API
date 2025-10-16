package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectTODB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env: ", err)
		return
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Failed to connect to Database ma boi: ", err)
		return
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Failed to Ping Database! wtf: ", err)
		return
	}

	fmt.Println("[*_*] Connected to database hehehe!")
}
