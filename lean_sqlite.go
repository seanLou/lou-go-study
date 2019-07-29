package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, e := sql.Open("sqlite3", "./test_sqlite")
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	stmt, e := db.Prepare("INSERT INTO user(name, password, create_time) VALUES (?,?,?)")
	if e != nil {
		log.Fatal(e)
	}
	result, e := stmt.Exec("louguanyang", "123456", time.Now().Format("yyyy-MM-dd"))
	if e != nil {
		log.Fatal(e)
	}
	lastInsertId, e := result.LastInsertId()
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(lastInsertId)
}
