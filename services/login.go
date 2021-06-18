package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Agent struct {
	username string
	password string
}

var db *sql.DB

func main() {
	fmt.Println("login")

	db, err := sql.Open("mysql", "root:TsP3#86b29@/dac_01")
	if err != nil {
		fmt.Println("unable to connect to database")
	}

	username := CheckUsername(db)
	password := CheckPassword(db, username)
	fmt.Println("username:", " "+username+"password:", password)
}

func CheckUsername(db *sql.DB) string {
	advisor := Agent{}
	var username string
	fmt.Println("enter your username")
	fmt.Scanln(&username)

	query := "select agentUsername from agentLogin where agentUsername = (?)"

	agentUsername, err := db.Query(query, username)
	if err != nil {
		fmt.Println("error with query, please check")
	}

	//run through the rows as requested
	for agentUsername.Next() {
		err = agentUsername.Scan(&advisor.username)
	}

	if username == advisor.username {
		fmt.Println(advisor.username)
	} else {
		fmt.Println("username does not exist")
	}
	return advisor.username
}

func CheckPassword(db *sql.DB, username string) string {
	advisorP := Agent{}
	var password string
	fmt.Println("enter your password")
	fmt.Scanln(&password)

	query := "select agentPassword from agentLogin where agentUsername = (?)"

	agent_password, err := db.Query(query, username)
	if err != nil {
		fmt.Println("error with query, please check")
	}
	//run through the rows as requested
	for agent_password.Next() {
		err = agent_password.Scan(&advisorP.password)
	}

	if password == advisorP.password {
		fmt.Println("correct password")
	} else if password != advisorP.password {
		fmt.Println("password is incorrect, try again")
	}
	defer db.Close()
	return advisorP.password
}

func handleRequest(w http.ResponseWriter, r http.Request) {

}
