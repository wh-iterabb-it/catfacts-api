package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Repository ...
type Repository struct{}

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func NewRepositoryPostgres() {
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// tests an initial connection to see if possible
	err = db.Ping()
	if err != nil {
		// initial connection failed
		fmt.Println("Failure on connection!")
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func NewRepositoryMysql() {
	db := sqlx.MustConnect("mysql", "root:root@tcp(127.0.0.1:3306)/sqlx_test")
	if err != nil {
		panic(err)
	}

}
