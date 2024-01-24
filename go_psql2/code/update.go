package db

import (
	"database/sql"
	"postgresqlDatabase/go_psql2/models"
)

// this function updates Student from table on database
// this function increments Student's age by 1
func UpdateStudent(db *sql.DB, student *models.Student) (*models.Student, error) {
	// query of update
	query := `
	UPDATE 
		students 
	SET 
		age = age+1 
	WHERE 
		id = $1 
	RETURNING 
		id, 
		first_name, 
		last_name, 
		age`

	// run the query
	StudentRows := db.QueryRow(query, student.Id)

	// copy query result
	var NewStudent models.Student
	scanErr := StudentRows.Scan(
		&NewStudent.Id,
		&NewStudent.FirstName,
		&NewStudent.LastName,
		&NewStudent.Age,
	)
	NewStudent.CourseId = student.CourseId

	// copy error checking
	if scanErr != nil {
		return &models.Student{}, scanErr
	}

	return &NewStudent, nil
}

// this function updates Course from table on database
// this function increases the cost of the Course by 10%
func UpdateCourse(db *sql.DB, course *models.Course) (*models.Course, error) {
	// query of update
	query := `
	UPDATE 
		courses 
	SET 
		price = price + price*0.1 
	WHERE 
		id = $1 
	RETURNING 
		id, 
		name, 
		teacher, 
		price`

	// run the query
	CourseRows := db.QueryRow(query, course.Id)

	// copy query result
	var NewCourse models.Course
	scanErr := CourseRows.Scan(
		&NewCourse.Id,
		&NewCourse.Name,
		&NewCourse.Teacher,
		&NewCourse.Price,
	)
	NewCourse.StudentId = course.StudentId

	// copy error checking
	if scanErr != nil {
		return &models.Course{}, scanErr
	}

	return &NewCourse, nil
}