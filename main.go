package  main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"time"
)

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456aA@@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connect successfully")
	return db
}

func main() {
	db := connectDB()
	defer  db.Close()
	id := uuid.New().String()
	firstName := "Quan"
	lastName := "Nguyen"
	middleName := "Sy"
	createdAt := time.Now()
	updatedAt := time.Now()
	insForm, err := db.Prepare("INSERT INTO user(id, first_name, middle_name, last_name, created_at, updated_at) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(id, firstName, lastName, middleName, createdAt, updatedAt)
	fmt.Println("Insert succesfully")
}