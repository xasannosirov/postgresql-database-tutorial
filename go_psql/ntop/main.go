package main

import (
	"database/sql"
	"fmt"

	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
)

//Student nomli struct 
type Students struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
}

// Courses nomli stuct
type Courses struct {
	Id      int
	Name    string
	Price   int
	Teacher string
}

func main() {

	// databasega connect qilish
	connection := "user=newuser password=1234 dbname=newdb sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// select query yozish (bir kursdagi talabalar)
	rows, err := db.Query(`
	SELECT 
    	s.first_name AS StudentFirstName,
    	s.last_name AS StudentLastName,
    	s.age AS Age,
    	c.name AS CourseNmae,
    	c.teacher AS CourseTeacher,
    	c.price AS CoursePrice
	FROM
	    Students s 
	INNER JOIN
	    StudentCourse sc ON s.id = sc.student_id
	INNER JOIN 
	    Courses c ON c.id = sc.course_id
	WHERE
	    c.name = $1`, "SMM")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// kursdagi talabalarni chiqarish
	for rows.Next() {
		s := Students{}
		c := Courses{}
		err := rows.Scan(
			&s.FirstName,
			&s.LastName,
			&s.Age,
			&c.Name,
			&c.Teacher,
			&c.Price,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		pp.Println(s)
		pp.Println(c)
	}

	// select query yozish (bir kursdagi talabalar)
	rows, err = db.Query(`
	SELECT 
    	s.first_name AS StudentFirstName,
    	s.last_name AS StudentLastName,
    	s.age AS Age,
    	c.name AS CourseNmae,
    	c.teacher AS CourseTeacher,
    	c.price AS CoursePrice
	FROM
	    Students s 
	INNER JOIN
	    StudentCourse sc ON s.id = sc.student_id
	INNER JOIN 
	    Courses c ON c.id = sc.course_id
	WHERE
    s.id = $1;`, 1)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// talabani kurslarini chiqarish
	for rows.Next() {
		s := Students{}
		c := Courses{}
		err := rows.Scan(
			&s.FirstName,
			&s.LastName,
			&s.Age,
			&c.Name,
			&c.Teacher,
			&c.Price,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		pp.Println(s)
		pp.Println(c)
	}
}
