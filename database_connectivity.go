package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func database_connectivity() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println("Error connecting to SQLite:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging SQLite database:", err)
		return
	}
	fmt.Println("Successfully connected to SQLite database")
}
