package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

// Books nomli struct
type Books struct {
	Id    int
	Name  string
	Price int
}

// Authors nomli struct
type Authors struct {
	Id        int
	LastName  string
	FirstName string
}

// BookAuthor nomli struct
type BookAuthor struct {
	Author_id int
	Book_id   int
}

func main() {

	//json ma'lumotlar
	reqBook := []byte(`{"name":"Book2","price":30}`)
	reqAuthor := []byte(`{"lastname":"Author2LastName","firstname":"Author2FirstName"}`)
	reqBookAuthor := []byte(`{"book_id":1, "author_id":2}`)

	//databasega connnect qilish
	connection := "user=newuser password=1234 dbname=newdb sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Json --> Books
	var book Books
	if err := json.Unmarshal(reqBook, &book); err != nil {
		panic(err)
	}

	//Bookni databasega insert qilish va return qilish
	var resBook Books
	sqlRowBook := db.QueryRow(`INSERT INTO books (name, price) VALUES ($1, $2) RETURNING name, price`, book.Name, book.Price)

	//return ma'lumotini Books structga olish
	if err := sqlRowBook.Scan(&resBook.Name, &resBook.Price); err != nil {
		panic(err)
	}

	fmt.Println("Book succesfully inserted with 'name', 'price':", resBook.Name, resBook.Price)

	//Json --> Authors
	var author Authors
	if err := json.Unmarshal(reqAuthor, &author); err != nil {
		panic(err)
	}

	//Authorni databasega insert qilish va return qilish
	var resAuthor Authors
	sqlRowAuthor := db.QueryRow(`INSERT INTO authors (last_name, first_name) VALUES ($1, $2) RETURNING last_name, first_name`, author.LastName, author.FirstName)

	//return ma'lumotini Authors structiga olish
	if err := sqlRowAuthor.Scan(&resAuthor.LastName, &resAuthor.FirstName); err != nil {
		panic(err)
	}

	fmt.Println("Author succesfully inserted with 'last_name', 'first_name':", resAuthor.LastName, resAuthor.FirstName)

	//Json --> BookAuthor
	var bookauthor BookAuthor
	if err := json.Unmarshal(reqBookAuthor, &bookauthor); err != nil {
		panic(err)
	}

	//BookAuthorni databasega insert qilish va return qilish
	var resBookAuthor BookAuthor
	sqlRowBookAuthor := db.QueryRow(`INSERT INTO bookauthor (book_id, author_id) VALUES ($1, $2) RETURNING book_id, author_id`, bookauthor.Book_id, bookauthor.Author_id)

	//return ma'lumotini BookAuthor structga olish
	if err := sqlRowBookAuthor.Scan(&resBookAuthor.Book_id, &resBookAuthor.Author_id); err != nil {
		panic(err)
	}

	fmt.Println("BookAuthor succesfully inserted with 'book_id', 'author_id':", resBookAuthor.Book_id, resBookAuthor.Author_id)

}
