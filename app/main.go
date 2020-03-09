package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Item is structure
type Item struct {
	ID    int    `db:"id"`
	Label string `db:"label"`
	Value string `db:"value"`
}

var schema string = "CREATE TABLE `items` (	  	`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,		  	`label` varchar(255) NOT NULL,		  	`value` varchar(255) NOT NULL		)"

func main() {
	conn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/story")
	if err != nil {
		panic(err)
	}
	conn.MustExec(schema)
	// res, err := conn.Exec("INSERT INTO items (name) VALUES(\"Peter\")")
	// if err != nil {
	// 	panic(err)
	// }
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Created user with id:%d", id)
	// _, err = conn.Exec("UPDATE items set name=\"John\" where id=?", id)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = conn.Exec("DELETE FROM items where id=?", id)
	// if err != nil {
	// 	panic(err)
	// }

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/items" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		var users []Item
		err = conn.Get(&users, "select * from items")
		if err != nil {
			panic(err)
		}

		if r.Method == "GET" {
			value, _ := json.Marshal(users)
			w.Write(value)
		}

	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
