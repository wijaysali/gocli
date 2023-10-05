package handler

import (
	"basic/entity"
	"database/sql"
)

func ListStudents(db *sql.DB) ([]entity.Student, error) {
	rows, err := db.Query("SELECT id, name, email FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []entity.Student
	for rows.Next() {
		var s entity.Student
		err := rows.Scan(&s.ID, &s.Name, &s.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}
