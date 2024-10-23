package main

import (
	"database/sql"
	"fagents/db"
	"fagents/tg"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
	//_ "github.com/mattn/go-sqlite3"
)

func main() {
	var (
		err  error
		path string
	)

	if len(os.Args) < 2 {
		panic("Bot token not specified")
	}

	tg.FagentsBotToken = os.Args[1]
	path, err = os.Executable()
	if err != nil {
		panic(err)
	}

	path = filepath.Dir(path)

	db.DBConn, err = sql.Open("sqlite", path+"/db/fagents.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.DBConn.Close()

	queryText := "SELECT * FROM fagents WHERE fullName LIKE ? OR inn LIKE ? OR members LIKE ?"
	db.Statement, err = db.DBConn.Prepare(queryText)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Statement.Close()

	tg.InitBot()
}
