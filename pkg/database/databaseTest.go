package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//*USE THIS AS A TEST FOR THE DATABASE IE ALL CRUD OPERATIONS

var db *sql.DB

var username string = "root"
var password string = "TsP3#86b29"
var hostname string = "127.0.0.1:3306"
var dbname string = "dac_01"

type Test struct {
	Id      int
	Message string
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "username:password@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}
	return db
}

func Create(db *sql.DB) {
	query := `insert into test values ((?), (?))`
	var id int
	var message string
	fmt.Scanln(&id, &message)
	_, err := db.Exec(query, id, message)
	if err != nil {
		fmt.Println("unable to insert records")
	}
}

func Read(db *sql.DB) {
	fmt.Println("enter a query without quotes")
	//have to use scanner as fmt.Scanln(&*) doesn't like spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	new := scanner.Text()
	//query := "select * from test"
	rows, err := db.Query(new)
	if err != nil {
		fmt.Println("error retrieving records")
	}

	var test = Test{}
	//all the columns in the db must be requested otherwise
	//the results are returned as 0 - boooooo!!
	for rows.Next() {
		err = rows.Scan(&test.Id, &test.Message)
		fmt.Println(test)
	}
}

func Update(db *sql.DB) {

}

func Delete(db *sql.DB) {
	//query := `delete from test where id = $1`
	//_, err := db.Exec("delete from test where id = 4")
	//if err != nil {
	//fmt.Println("unable to delete record")
	//}
}

func main() {

	db, err := sql.Open("mysql", "root:TsP3#86b29@/dac_01")
	if err != nil {
		fmt.Println("error connecting to:", dbname)
	}

	var choice int

	fmt.Println("what operations would you like to test?")
	fmt.Println("1: create, 2: read, 3: update 4: delete")
	fmt.Scanln(&choice)
	if choice == 1 {
		Create(db)
	} else if choice == 2 {
		Read(db)
	} else if choice == 3 {
		Update(db)
	} else if choice == 4 {
		Delete(db)
	}

	defer db.Close()
}
