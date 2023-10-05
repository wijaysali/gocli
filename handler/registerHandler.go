package handler

import (
	"database/sql"
	"basic/entity"
)

func RegisterStudent(db *sql.DB, student entity.Student)error {
	_, err := db.Exec("INSERT INTO students (name, email) VALUES (?,?)", student.Name, student.Email)
	return err
}