package main

import (
	"database/sql"

	//	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// const (
// 	DB_USER     = "postgres"
// 	DB_PASSWORD = "postgres"
// 	DB_NAME     = "lumoswg"
// )

// func DbConnection() (*sql.DB, error) {
// 	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
// 		DB_USER, DB_PASSWORD, DB_NAME)
// 	db, err := sql.Open("postgres", dbinfo)
// 	CheckErr(err)
// 	//defer db.Close()
// 	return db, err
// }

func SQLiteConn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./pushnotification.db")
	//checkErr(err)
	return db, err
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
