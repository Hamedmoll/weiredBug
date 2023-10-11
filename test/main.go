package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:shopstoreRoo7t0lk2o20@tcp(localhost:4444)/shopstore_db")
	fmt.Printf("error %v\n:", err)
	var version string
	err2 := db.QueryRow("select version()").Scan(&version)

	fmt.Printf("error 2 %v\n%s", err2, version)
}
