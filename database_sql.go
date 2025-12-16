package main

import (
	"fmt"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func database_sql() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: ", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Connected to the database successfully!")
}