package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

var ErrNoMatch = fmt.Errorf("no matching record")

func Initialize(host string, dbName string) (Database, error) {
	db := Database{}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, PORT,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		dbName,
	)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Connected to database successfully")
	return db, nil
}
