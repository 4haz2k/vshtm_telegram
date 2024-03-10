package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type schedule struct {
	Id        int
	Notified  bool
	Subject   string
	Theme     string
	Link      string
	CreatedAt sql.NullTime
	Teacher   string
	Building  sql.NullString
}

var db *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("environment file not found")
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Panic(err)
	}
}

func LogMessage(chatId int64, message string) {
	_, err := db.Exec("INSERT INTO log (chat_id, message, created_at, updated_at) VALUES (?, ?, NOW(), NOW())", strconv.Itoa(int(chatId)), message)
	if err != nil {
		log.Fatal(err)
	}
}

func SubscribeUser(chatId int64) {
	_, err := db.Exec("INSERT INTO participants (chat_id, subscribed) VALUES (?, true) ON DUPLICATE KEY UPDATE subscribed = true;", strconv.Itoa(int(chatId)))
	if err != nil {
		log.Fatal(err)
	}
}

func UnsubscribeUser(chatId int64) {
	_, err := db.Exec("INSERT INTO participants (chat_id, subscribed) VALUES (?, false) ON DUPLICATE KEY UPDATE subscribed = false;", strconv.Itoa(int(chatId)))
	if err != nil {
		log.Fatal(err)
	}
}

func GetSchedule() []schedule {
	startDate := time.Now().Local().Add(time.Hour * 3).Format("2006-01-02 15:04:05")
	endDate := time.Now().Local().Add(time.Hour * 171).Format("2006-01-02 15:04:05") // week + 3 hours

	rows, err := db.Query("SELECT * FROM schedule WHERE created_at BETWEEN ? AND ?;", startDate, endDate)
	if err != nil {
		log.Fatal(err)
	}

	var list []schedule

	for rows.Next() {
		s := schedule{}
		err := rows.Scan(&s.Id, &s.Notified, &s.Subject, &s.Theme, &s.Link, &s.CreatedAt, &s.Teacher, &s.Building)
		if err != nil {
			fmt.Println(err)
			continue
		}

		list = append(list, s)
	}

	return list
}
