package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var username string = "root"
var password string = "TsP3#86b29"
var hostname string = "127.0.0.1:3306"
var dbname string = "dac_01"

type Test struct {
	Id      int
	Message string
}

func Connect() {
	db, err := sql.Open("mysql", "root:TsP3#86b29@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}
}

func TestConnection(db *sql.DB) {
	var test = Test{}
	rows, err := db.Query("select * from test")
	if err != nil {
		fmt.Println("error with query, please check")
	}
	for rows.Next() {
		err = rows.Scan(&test.Id, &test.Message)
		fmt.Println(test)
	}
}

//this is intended as a default query handler where a dedicated function is overkill
func Query(db *sql.DB, query string) {
	//can a default struct be created to serve multiple types
	var test = Test{}
	//receive the query string
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("error with query, please check")
	}
	//run through the rows as requested
	for rows.Next() {
		err = rows.Scan(&test.Id, &test.Message)
		fmt.Println(test)
	}
}

func main() {
	Connect()
}
