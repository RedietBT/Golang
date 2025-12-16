package main
import (
	"fmt"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectMySQL() (*sql.DB, error) {
	dsn := "username:password@tcp(localhost:3306)/database_name"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func readRecords(db *sql.DB) {
	query := "SELECT id, name, age FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Failed to read records: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string 
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal("Failed to scan record: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

func UpdateRecord(db *sql.DB) {
	query := "UPDATE users SET  age = ? WHERE name = ?"
	result, err := db.Exec(query, 35, "John Doe")	
	if err != nil {
		log.Fatalf("Failed to update record: %v", err)
	}
	fmt.Printf("Updated row: %d\n", rowAffected)
}

func main() {
	db, err := connectMySQL()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	createRecord(db)
	readRecords(db)
	UpdateRecord(db)
	deleteRecord(db)
}