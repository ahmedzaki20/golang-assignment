package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB
func InitDB()  {
	var err error
	DB ,err=sql.Open("sqlite3","api.db")
	if err !=nil{
		panic(err)
	}
DB.SetMaxIdleConns(5)
DB.SetMaxOpenConns(10)

createTables()
}

func createTables()   {
	sqlStmt := `CREATE TABLE IF NOT EXISTS animals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER,
		breed TEXT)`
	_,err:=DB.Exec(sqlStmt)
	if err!=nil{
		panic(err)
	}
	return
}