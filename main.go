package main

import (
	"database/sql"
	"fagents/db"
	"fagents/types"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	conn        *sql.DB
	fagentsList []types.Fagent
	searchText  string
)

func main() {
	conn, err := sql.Open("sqlite3", "./db/fagents.sqlite")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Print("Условие поиска: ")
	fmt.Scanln(&searchText)
	fmt.Println("Поиск по условию:", searchText)

	fagentsList, err := db.FagentsList(conn, searchText)
	if err != nil {
		fmt.Println("Ошибка при поиске", err)
		os.Exit(1)
	}

	sendToTg(fagentsList)

}

func sendToTg(faList []types.Fagent) {
	if len(faList) == 0 {
		fmt.Println("Ничего не найдено")
	} else {
		for i, v := range faList {
			fmt.Println(i, v.Inn, v.FullName)
		}
	}
}
