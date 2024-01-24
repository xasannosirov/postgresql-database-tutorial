package db

import (
	"database/sql"
	"postgresqlDatabase/go_psql2/models"
)

// this func gets Student with all Courses
func GetStudent(db *sql.DB, studentId int) (*models.Student, error) {
	// get query
	queryUser := `
	SELECT 
		id, 
		first_name, 
		last_name, 
		age 
	FROM 
		students 
	WHERE 
		id = $1`

	// run the query
	StudentRow := db.QueryRow(queryUser, studentId)

	// copy query result
	var NewStudent models.Student
	sErr := StudentRow.Scan(
		&NewStudent.Id,
		&NewStudent.FirstName,
		&NewStudent.LastName,
		&NewStudent.Age,
	)

	// scan error checking
	if sErr != nil {
		return &models.Student{}, sErr
	}

	// get courses of student query
	queryUserCourse := `
	SELECT 
		course_id 
	FROM 
		studentcourse 
	WHERE 
		student_id = $1`

	// run the query
	respCourseIds, cErr := db.Query(queryUserCourse, NewStudent.Id)

	// checking error
	if cErr != nil {
		return &models.Student{}, cErr
	}

	// close rows
	defer respCourseIds.Close()

	// get all ids of course
	for respCourseIds.Next() {
		var newId int
		sErr := respCourseIds.Scan(&newId)
		if sErr != nil {
			return &models.Student{}, sErr
		}
		NewStudent.CourseId = append(NewStudent.CourseId, newId)
	}

	return &NewStudent, nil
}

// this func gets Course with all Students
func GetCourse(db *sql.DB, courseId int) (*models.Course, error) {
	// get query
	queryCourse := `
	SELECT 
		id, 
		name, 
		teacher, 
		price 
	FROM 
		courses 
	WHERE 
		id = $1`

	// run the query
	Courserow := db.QueryRow(queryCourse, courseId)

	// copy result query
	var NewCourse models.Course
	sErr := Courserow.Scan(
		&NewCourse.Id,
		&NewCourse.Name,
		&NewCourse.Teacher,
		&NewCourse.Price,
	)

	// scan error checking
	if sErr != nil {
		return &models.Course{}, sErr
	}

	// get all ids of student query
	queryUserCourse := `
	SELECT 
		student_id
	FROM 
		studentcourse 
	WHERE 
		course_id = $1`

	// run the query
	respStudentIds, cErr := db.Query(queryUserCourse, NewCourse.Id)

	// checking error
	if cErr != nil {
		return &models.Course{}, cErr
	}

	// close rows
	defer respStudentIds.Close()

	// get all ids of student
	for respStudentIds.Next() {
		var newId int
		sErr := respStudentIds.Scan(&newId)
		if sErr != nil {
			return &models.Course{}, sErr
		}
		NewCourse.StudentId = append(NewCourse.StudentId, newId)
	}

	return &NewCourse, nil
}