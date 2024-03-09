package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
)

var db *sql.DB

func init() {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USERNAME"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_DATABASE"),
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panic(err)
	}
}

func LogMessage(chatId int64, message string) {
	_, err := db.Exec("INSERT INTO log (chat_id, message) VALUES (?, ?)", strconv.Itoa(int(chatId)), message)
	if err != nil {
		log.Fatal(err)
	}
}
