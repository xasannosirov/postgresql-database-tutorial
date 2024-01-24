package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Books nomli sturct
type Books struct {
	Id    int
	Name  string
	Price int
}

func main() {
	// databasega connect qilish
	connection := "user=newuser password=1234 dbname=newdb sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// select query yozish
	rows, err := db.Query(`
	SELECT 
		id, 
		name, 
		price 
	FROM 
		books 
	WHERE 
		price = $1`, 20)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// query natijasini structga olish
	book := []Books{}
	for rows.Next() {
		b := Books{}
		err := rows.Scan(
			&b.Id,
			&b.Name,
			&b.Price,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		book = append(book, b)
	}

	// natijani ko'rish
	for _, b := range book {
		fmt.Printf("Id: %d\nName: %s\nPrice: %d\n\n", b.Id, b.Name, b.Price)
	}
}
