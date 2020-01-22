package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)
//sql file
//CREATE TABLE users (
//id SERIAL PRIMARY KEY,
//age INT,
//first_name TEXT,
//last_name TEXT,
//email TEXT UNIQUE NOT NULL
//);
func main()  {
	db, err := sql.Open("postgres","user=postgres password=postgres host=localhost port=5432 dbname=golang sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect successfully")
	//insert
	//sqlStatement := "INSERT INTO users1 (age, email, first_name, last_name) VALUES ($1, $2, $3, $4)"
	//for i := 1; i < 10000001; i++ {
	//	s := strconv.Itoa(i)
	//	result, err := db.Exec(sqlStatement, i, "rpa" + s +"@topica.edu.vn", "Sơn", "Trần Minh")
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	numberInsert, err := result.RowsAffected()
	//	fmt.Println(numberInsert)
	//}

	//lay id sau khi insert
	sqlStatement := `INSERT INTO users1 (age, email, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 30, "thaottp7@topica.edu.vn", "thảo", "trần thị phương").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record Id is:", id)
}
