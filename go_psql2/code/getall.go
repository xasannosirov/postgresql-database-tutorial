package db

import (
	"database/sql"
	"postgresqlDatabase/go_psql2/models"
)

// this func gets all Students with all Courses
func GetAllStudent(db *sql.DB) ([]*models.Student, error) {
	var allStudents []*models.Student
	// get all student with id, first_name, last_name, age
	query := `
	SELECT 
		id, 
		first_name, 
		last_name, 
		age
	FROM
		students`

	// run the query
	AllRowStudents, aErr := db.Query(query)

	// query error checking
	if aErr != nil {
		return []*models.Student{}, aErr
	}

	// all students append a slice
	for AllRowStudents.Next() {
		// get a new student and copy to new struct
		var OneStudent models.Student
		serr := AllRowStudents.Scan(
			&OneStudent.Id,
			&OneStudent.FirstName,
			&OneStudent.LastName,
			&OneStudent.Age,
		)

		// scan error checking
		if serr != nil {
			return []*models.Student{}, serr
		}
		// get this student course ids
		queryIds := `
		SELECT 
			course_id
		FROM
			studentcourse
		WHERE
			student_id = $1`
			
		// run the query
		rowsIds, rErr := db.Query(queryIds, OneStudent.Id)

		// query error checking
		if rErr != nil {
			return []*models.Student{}, rErr
		}

		// get course ids
		var ids []int
		for rowsIds.Next() {
			var id int
			sErr := rowsIds.Scan(&id)
			if sErr != nil {
				return []*models.Student{}, sErr
			}
			ids = append(ids, id)
		}
		// ids copy to struct
		OneStudent.CourseId = ids
		// student append to student slice
		allStudents = append(allStudents, &OneStudent)
	}

	return allStudents, nil
}

// this func gets all Courses with all Students
func GetAllCourse(db *sql.DB) ([]*models.Course, error) {
	var allCourses []*models.Course
	// get all course with id, name, teacher, price
	query := `
	SELECT 
		id, 
		name, 
		teacher, 
		price
	FROM
		courses`

	// run the query
	AllRowCourses, aErr := db.Query(query)

	// query error checking
	if aErr != nil {
		return []*models.Course{}, aErr
	}

	// get all course
	for AllRowCourses.Next() {
		// get one course
		var OneCourse models.Course
		serr := AllRowCourses.Scan(
			&OneCourse.Id,
			&OneCourse.Name,
			&OneCourse.Teacher,
			&OneCourse.Price,
		)

		// scan error checking
		if serr != nil {
			return []*models.Course{}, serr
		}

		// get course with student ids
		queryIds := `
		SELECT 
		student_id
		FROM
		studentcourse
		WHERE
		course_id = $1`

		// run the query
		rowsIds, rErr := db.Query(queryIds, OneCourse.Id)

		// query error checking
		if rErr != nil {
			return []*models.Course{}, rErr
		}

		// get all student_ids
		var ids []int
		for rowsIds.Next() {
			var id int
			sErr := rowsIds.Scan(&id)
			if sErr != nil {
				return []*models.Course{}, sErr
			}
			ids = append(ids, id)
		}

		// copy student ids
		OneCourse.StudentId = ids

		// appedn course to course slice
		allCourses = append(allCourses, &OneCourse)
	}

	return allCourses, nil
}