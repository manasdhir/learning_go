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
	defer db.Close()
	db.Exec("Create database if not exists test")
	db.Exec("use test")
	db.Exec("create table User(ind int auto_increment primary key, name varchar(100))")
	err1 := db.Ping()
	if err1 != nil {
		fmt.Println("Error connecting to db")
	}
	fmt.Println("connected to mysql")
	createuser("amarqqqaaaaa", db)
	readusers(db)
	updateuser(1, "amarbhadwa", db)
	deleteuser(6, db)
	readusers(db)
}

func createuser(name string, con *sql.DB) {
	query := "Insert into User(name) VALUES(?)"
	res, err := con.Exec(query, name)
	if err != nil {
		fmt.Println("error")
	}
	id, _ := res.LastInsertId()
	fmt.Println(id)
}
func readusers(con *sql.DB) {
	query := "select ind, name from User"
	res, err := con.Query(query)
	if err != nil {
		fmt.Println("error")
	}
	defer res.Close()
	var id int
	var name string
	for res.Next() {
		res.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

func updateuser(id int, name string, con *sql.DB) {
	query := "update User set name= ? where ind=?"
	res, err := con.Exec(query, name, id)
	if err != nil {
		fmt.Println("error while updating")
	}
	ra, _ := res.RowsAffected()
	if ra == 0 {
		fmt.Println("not found")
	} else {
		fmt.Println("no. of rows effected", ra)
	}
}
func deleteuser(id int, con *sql.DB) {
	query := "delete from User where ind=?"
	res, err := con.Exec(query, id)
	if err != nil {
		fmt.Println("error removing user")
	}
	ras, _ := res.RowsAffected()
	if ras == 0 {
		fmt.Println("user does not exist")
	}
	fmt.Println(ras, "rows effected")
}
