package main

import (
	"database/sql"
	"fagents/db"
	"fagents/tg"
	"fmt"

	_ "modernc.org/sqlite"
	//_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn, err := sql.Open("sqlite", "./db/fagents.sqlite")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	db.DBConn = conn

	queryText := "SELECT * FROM fagents WHERE fullName LIKE ? OR inn LIKE ? OR members LIKE ?"
	db.Statement, err = db.DBConn.Prepare(queryText)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Statement.Close()

	tg.InitBot()
}
