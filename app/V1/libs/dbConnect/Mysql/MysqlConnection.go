package MysqlConnection

import (
    "database/sql"
    "fmt"
)

func DbConn() (db *sql.DB) {
    /* dbDriver := "mysql"
    dbUser := "root"
	dbPass := "12345"
	dbHost := "localhost" 
	dbPort := "3306" 
    dbName := "trucklagbe_db" */
	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+"/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/p4tl?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(db.Ping())
    return db
}
