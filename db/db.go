package db

import (
	"database/sql"
	"fagents/types"
	"fmt"
	"strings"
)

func FagentsList(db *sql.DB, searchText string) ([]types.Fagent, error) {
	queryText := "SELECT * FROM fagents WHERE upper(fullName) LIKE '%" + strings.ToUpper(searchText) + "%' OR inn LIKE '%" + searchText + "%'"
	//queryText := "SELECT id, upper(fullName), dob, ogrn, inn, regNum, snils, address, resources, members, law, dateIn, datePubl, dateOut FROM fagents WHERE upper(fullName) LIKE '%" + strings.ToUpper(searchText) + "%' OR inn LIKE '%" + searchText + "%'"
	fmt.Println(queryText)
	rows, err := db.Query(queryText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []types.Fagent{}
	for rows.Next() {
		i := types.Fagent{}
		err = rows.Scan(&i.Id, &i.FullName, &i.Dob, &i.Ogrn, &i.Inn, &i.RegNum, &i.Snils, &i.Address, &i.Resources, &i.Members, &i.Law, &i.DateIn, &i.DatePubl, &i.DateOut)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}

	return data, nil
}
