package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID        int    `json:"id"`
	LastName  string `json:"lname"`
	FirstName string `json:"fname"`
	Age       int    `json:"age"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:Root@123@tcp(localhost:3306)/sakila")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("insert into persons values(6,'A','S',100)")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	// Execute the query
	results, err := db.Query("SELECT ID, age FROM persons")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("get all data")

	for results.Next() {
		var p Person
		// for each row, scan the result into our tag composite object
		err = results.Scan(&p.ID, &p.Age)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Println(p.ID, p.Age)
	}

	fmt.Println("get specific data")
	var ps Person
	err = db.QueryRow("SELECT ID, Age FROM persons where id = ?", 2).Scan(&ps.ID, &ps.Age)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(ps.ID, ps.Age)
}
