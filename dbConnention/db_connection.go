package dbConnention

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

type DBConnection interface {
	Open() error
	Close() error
	GetDb() *sql.DB
}

type DBConnectionStruct struct {
	db *sql.DB
}

func (dbs *DBConnectionStruct) Open() error {
	log.Println("Open connection to DB")
	tempDb, err := sql.Open("mssql", "server=localhost;user id=sa;password=123456;database=TodoDB")

	if err != nil {
		log.Println(err)
		return err
	}

	dbs.db = tempDb
	log.Println("Connected to DB")
	return err
}

func (dbs *DBConnectionStruct) Close() error {
	err := dbs.db.Close()

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Close connection to db")
	return err
}

func (dbs *DBConnectionStruct) GetDb() *sql.DB {
	return dbs.db
}