package main

import (
	"database/sql"
	code "postgresqlDatabase/go_psql2/code"
	models "postgresqlDatabase/go_psql2/models"

	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
)

func main() {
	// connect to database
	connection := "user=newuser password=1234 dbname=newdb sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Create Course object
	course := models.Course{
		Name:      ".Net",
		Teacher:   "Muhammadkarim Tukhtaboev",
		Price:     2200000,
		StudentId: []int{4, 3},
	}

	// Create a new Course with all Students
	cCourse, cErr := code.CreateCourse(db, &course)
	if cErr != nil {
		panic(cErr)
	}
	pp.Println(cCourse)

	// // Updates Course
	uCourse, uErr := code.UpdateCourse(db, cCourse)
	if uErr != nil {
		panic(err)
	}
	pp.Println(uCourse)

	// // Gets a Course with all Students
	gCourse, gerr := code.GetCourse(db, uCourse.Id)
	if gerr != nil {
		panic(err)
	}
	pp.Println(gCourse)

	// // Delete a Course with  all Students
	if dErr := code.DeleteCourse(db, gCourse); dErr != nil {
		panic(dErr)
	}
	pp.Println("Deleted Course")

	// // Get all Courses with all Students
	var allCourse []*models.Course
	allCourse, allerr := code.GetAllCourse(db)
	if allerr != nil {
		panic(allerr)
	}
	pp.Println(allCourse)

	// // Create sturct object
	student := models.Student{
		Id:        2,
		FirstName: "Ahrorbek",
		LastName:  "Olimjonov",
		Age:       21,
		CourseId:  []int{2, 3},
	}

	// // Create a new Student with all Courses
	CStudent, cErr := code.CreateStudent(db, &student)
	if cErr != nil {
		panic(cErr)
	}
	pp.Println(CStudent)

	// // Updates Student
	UStudent, uErr := code.UpdateStudent(db, CStudent)
	if uErr != nil {
		panic(uErr)
	}
	pp.Println(UStudent)

	// // Get a Student with all Coruses
	GStudent, gErr := code.GetStudent(db, UStudent.Id)
	if gErr != nil {
		panic(gErr)
	}
	pp.Println(GStudent)

	// // // Delete a Student with all Courses
	if err := code.DeleteStudent(db, GStudent); err != nil {
		panic(err)
	}
	pp.Println("Deleted Student")

	// // Get all Students with all Courses
	var allStudents []*models.Student
	allStudents, allErr := code.GetAllStudent(db)
	if allErr != nil {
		panic(allErr)
	}
	pp.Println(allStudents)

	pp.Println("Succesfully completed all tasks")
}