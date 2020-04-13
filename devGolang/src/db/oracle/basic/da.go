package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/rana/ora.v4"
)

func main() {
	db, err := sql.Open("ora", "ltuser/lt2005@192.168.2.57:1522/TESTLTDB")
	if err != nil {
		fmt.Println("Connection error", err)
	}
	defer db.Close()

	qry := "SELECT user FROM dual"
	fmt.Println("Running query:", qry)
	rows, err := db.Query(qry)

	if err != nil {
		fmt.Println("Error:", err)
	}

	defer rows.Close()
	var user string
	for rows.Next() {
		if err = rows.Scan(&user); err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println(user)
	}
}
