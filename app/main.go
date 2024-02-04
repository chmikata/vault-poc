package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	db := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
}
