package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "godb"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("Cannot Connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id BIGSERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TIMESTAMP NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create event table")
	}
}
