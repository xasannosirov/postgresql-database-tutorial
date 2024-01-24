package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	// databasega connect qilish
	connection := "user=newuser password=1234 dbname=newdb sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// id = 2 bo'lgan kitobni o'chirish
	result, err := db.Exec("DELETE FROM Books WHERE id = $1", 2)
	if err != nil {
		panic(err)
	}
	delLine, delErr := result.RowsAffected()
	if delErr != nil {
		panic(err)
	}
	fmt.Println(delLine) //o'chrilgan qatorlar soni
}
