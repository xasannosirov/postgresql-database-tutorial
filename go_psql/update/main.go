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

	// id = 1 bo'lgan kitobni narxini o'zgartirish
	result, err := db.Exec(`UPDATE Books SET price = $1 where id = $2`, 25, 1)
	if err != nil {
		panic(err)
	}
	newLine, uptErr := result.RowsAffected() 
	if uptErr != nil {
		panic(uptErr)
	}
	fmt.Println(newLine) //yangilangan qatorlar soni
}
