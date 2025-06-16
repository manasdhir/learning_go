package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:rootpass@tcp(localhost:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("unable to open database")
	}
	db.Exec("Create database if not exists test")
	err1 := db.Ping()
	if err1 != nil {
		fmt.Println("Error connecting to db")
	}
	fmt.Println("connected to mysql")
	db.Close()
}
