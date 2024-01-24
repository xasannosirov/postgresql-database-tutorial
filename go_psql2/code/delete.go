package db

import (
	"database/sql"
	"postgresqlDatabase/go_psql2/models"
)

// this func deleted Student from table on database
func DeleteStudent(db *sql.DB, student *models.Student) error {
	// deletion request
	query := `
		DELETE FROM 
			studentcourse 
		WHERE 
			student_id = $1`

	// run the query
	_, execErr := db.Exec(query, student.Id)

	// exec error checking
	if execErr != nil {
		return execErr
	}
	student.CourseId = []int{}

	// deletion request
	queryDel := `
	DELETE FROM 
		students 
	WHERE 
		id = $1`

	// run the query
	_, err := db.Exec(queryDel, student.Id)

	// exec error checking
	if err != nil {
		return err
	}

	return nil
}

// this func deleted Course from table on database
func DeleteCourse(db *sql.DB, course *models.Course) error {
	// deletion request
	query := `
		DELETE FROM 
			studentcourse 
		WHERE 
			course_id = $1`

	// exec error checking
	_, execErr := db.Exec(query, course.Id)

	// run the query
	if execErr != nil {
		return execErr
	}

	// deletion request
	queryDel := `
	DELETE FROM 
		courses 
	WHERE 
		id = $1`

	// run the query
	_, err := db.Exec(queryDel, course.Id)

	// exec error checking
	if err != nil {
		return err
	}

	return nil
}