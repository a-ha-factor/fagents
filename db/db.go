package db

import (
	"database/sql"
	"errors"
	"fagents/types"
	"fmt"
	"strings"
)

var (
	DBConn *sql.DB
	Statement *sql.Stmt
	err error
)

func FagentsList(searchText string) (types.FagentsList, error) {
	if len(searchText) < 4 || len(searchText) > 30 {
		return nil, errors.New("неверная длина запроса")
	}

	rows, err := Statement.Query("%"+strings.ToUpper(searchText)+"%", "%"+searchText+"%", "%"+strings.ToUpper(searchText)+"%")
	if err != nil {
		fmt.Println(err)
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
