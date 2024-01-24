package db

import (
	"database/sql"
	"postgresqlDatabase/go_psql2/models"
)

// this func inserts Student to database
func CreateStudent(db *sql.DB, student *models.Student) (*models.Student, error) {
	// query for create a new user
	query := `
	INSERT INTO students (first_name, last_name, age)
	VALUES ($1, $2, $3) RETURNING id, first_name, last_name, age`

	// run the query
	StudentRow := db.QueryRow(query, student.FirstName, student.LastName, student.Age)

	// copy result of the query
	var NewStudent models.Student
	scanErr := StudentRow.Scan(
		&NewStudent.Id,
		&NewStudent.FirstName,
		&NewStudent.LastName,
		&NewStudent.Age,
	)

	// scan error checking
	if scanErr != nil {
		return &models.Student{}, scanErr
	}

	// connected course with student
	for _, courseId := range student.CourseId {
		// connect query
		query := `
		INSERT INTO studentcourse (student_id, course_id)
		VALUES ($1, $2)`

		// run the query
		_, eErr := db.Exec(query, NewStudent.Id, courseId)

		// exec error checking
		if eErr != nil {
			return &models.Student{}, eErr
		}
		// copy course ids to Newstudent.CourseId
		NewStudent.CourseId = student.CourseId
	}

	return &NewStudent, nil
}

// this func insert Course to database
func CreateCourse(db *sql.DB, course *models.Course) (*models.Course, error) {
	// query for create a new course
	query := `
	INSERT INTO courses (name, teacher, price)
	VALUES ($1, $2, $3) RETURNING id, name, teacher, price`

	// run the query
	CourseRows := db.QueryRow(query, course.Name, course.Teacher, course.Price)

	// copy result of the query
	var NewCourse models.Course
	scanErr := CourseRows.Scan(
		&NewCourse.Id,
		&NewCourse.Name,
		&NewCourse.Teacher,
		&NewCourse.Price,
	)

	// scan error checking
	if scanErr != nil {
		return &models.Course{}, scanErr
	}

	// added student_id to course
	for _, studentId := range course.StudentId {
		// get all student_id query
		query := `
		INSERT INTO studentcourse (student_id, course_id)
		VALUES ($1, $2)`

		// run the query
		_, eErr := db.Exec(query, studentId, NewCourse.Id)

		// exec error checking
		if eErr != nil {
			return &models.Course{}, eErr
		}

		// copy student ids to NewCourse.StudentId
		NewCourse.StudentId = course.StudentId
	}

	return &NewCourse, nil
}