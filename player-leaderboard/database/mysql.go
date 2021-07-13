package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"player-leaderboard/types"
)

func ConnectoDatabaseExample() *sql.DB{

	fmt.Print("Starting database serve...")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/truck")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}

func InsertIntoDataabse(table string, truck types.Truck){
	db := ConnectoDatabaseExample()
	//Formulating insert string to database
	is := "INSERT INTO " + table + " VALUES ('" + truck.Melancia + "','" + truck.Gatinha + "','" + truck.Picole + "')"

	insert, err := db.Query(is)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}