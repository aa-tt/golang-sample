package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price,string"`
}

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	statement, _ := DB.Prepare("CREATE TABLE IF NOT EXISTS products (id TEXT PRIMARY KEY, name TEXT, price REAL)")
	statement.Exec()
	fmt.Println("Database initialized in memory")
}
