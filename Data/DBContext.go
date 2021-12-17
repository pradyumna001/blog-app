package data

import (
	"database/sql"
	"fmt"

	//Postgress Driver
	_ "github.com/lib/pq"
)

const (
	DB_Host = "127.0.0.1"
	DB_PORT = 5432
	DB_USER = "postgres"
	DB_PASS = "123456"
	DB_NAME = "blog-app-db"
)

var DB *sql.DB
var connectionError error

//Database Connection
func init() {
	connStr := "user=" + DB_USER + " dbname=" + DB_NAME + " password=" + DB_PASS + " host=" + DB_Host + " sslmode=disable"
	DB, connectionError = sql.Open("postgres", connStr)
	if connectionError != nil {
		fmt.Println("First Error while oprning")
		panic(connectionError)
	}
	connectionError := DB.Ping()
	if connectionError != nil {
		fmt.Println("Second Error while pinging")

		panic(connectionError)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
}
