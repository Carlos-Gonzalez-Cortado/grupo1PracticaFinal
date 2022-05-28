package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var db *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func init() {
	err := godotenv.Load("/grupo1Server/app/config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	connInfo := connection{
		Host:     os.Getenv("MARIADB_URL"),
		Port:     os.Getenv("MARIADB_PORT"),
		User:     os.Getenv("MARIADB_USER"),
		Password: os.Getenv("MARIADB_PASSWORD"),
		DBName:   os.Getenv("MARIADB_DB"),
	}

	db, err = sql.Open("mysql", connToString(connInfo))
	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB is open\n")
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB pinged successfully\n")
	}
}

func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)

}
