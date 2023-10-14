package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Starting")
	db, err := sql.Open("mysql", "root:shopstoreRoo7t0lk2o20@tcp(localhost:4444)/shopstore_db")
	if err != nil {
		fmt.Printf("error %v \n", err)
		return
	}

	var version string
	err2 := db.QueryRow("select version()").Scan(&version)
	if err2 != nil {
		fmt.Printf("error %v \n", err2)
		return
	}
	fmt.Printf("Version %s \n", version)
}
