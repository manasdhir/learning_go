package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:rootpass@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer db.Close()
	http.HandleFunc("/create", createproduct)
	http.HandleFunc("/get", getproduct)
	http.ListenAndServe(":8000", nil)
}

func createproduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var p Product
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil || p.Name == "" || p.Price < 0 {
			http.Error(w, "Invalid Fields", http.StatusBadRequest)
		}
		query := "insert into product(name,price) VALUES(?,?)"
		res, err1 := db.Exec(query, p.Name, p.Price)
		if err1 != nil {
			http.Error(w, "Error creating Product", http.StatusInternalServerError)
		}
		lastid, _ := res.LastInsertId()
		p.Id = int(lastid)
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(p)

	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
func getproduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rows, err := db.Query("select id, name, price from product;")
		if err != nil {
			http.Error(w, "Error fetching Product", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var productarr []Product
		for rows.Next() {
			var product Product
			err := rows.Scan(&product.Id, &product.Name, &product.Price)
			if err != nil {
				http.Error(w, "Error scanning Product", http.StatusInternalServerError)
				return
			}
			productarr = append(productarr, product)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productarr)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}}
}