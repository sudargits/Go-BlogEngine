package main

import (
	"database/sql"
	"fmt"
)

const database  = "blog"
const host  =  "localhost"
const port = "3306"
const db = "mysql"
const username = "root"
const password  = ""

func connectDB() *sql.DB{
	db, err := sql.Open(db, username+"@tcp("+host+":"+port+")/"+database)
	checkErr(err)
	return db
}
func checkErr(err error) {
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}

func queryForRows(query string) (*sql.Rows,error){
	var db = connectDB()
	defer db.Close()

	rows, err := db.Query(query)
	checkErr(err)
	defer rows.Close()
	return rows,err
}
