package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	connStr := "root:@tcp(localhost:3306)/basic"
	db, err := sql.Open("mysql", connStr)
	return db, err
}
