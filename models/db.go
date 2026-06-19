package models

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price,string"`
}

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	if _, err := DB.Exec("CREATE TABLE IF NOT EXISTS products (id TEXT PRIMARY KEY, name TEXT, price REAL)"); err != nil {
		return err
	}

	fmt.Println("Database initialized in memory")
	return nil
}
