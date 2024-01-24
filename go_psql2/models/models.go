package models

// Struct called Student
type Student struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	CourseId  []int
}

// Struct called Course
type Course struct {
	Id        int
	Name      string
	Teacher   string
	Price     int
	StudentId []int
}