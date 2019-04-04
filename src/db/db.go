package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"strings"
)

type Db struct {
	Connection *sql.DB
	Status     bool
	Err        error
}

func (db *Db) ConnectionToMysqlServer(dbc *Config) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%d)/", dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Port)
	db.Connection, _ = sql.Open("mysql", connectionString)
	db.Err = db.Connection.Ping()
	if db.Err != nil {
		log.Printf("error connection to mysql server: \n%s\n", db.Err)
		db.Status = false
		return
	}
	log.Printf("Success connection to MySQL server\n")
	db.Status = true
	db.Err = nil
}

func (db *Db) CreateDB(name string) bool {
	if db.Status == false && db.Err != nil {
		log.Printf("Err, not connected to MySQL server\n")
		return false
	}
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	_, err := db.Connection.Exec(query)
	if err != nil {
		log.Printf("error creating database \"%s\": \n%s\n", name, err)
		return false
	}
	log.Printf("Success create DB: '%s'\n", name)
	return true
}

func (db *Db) ConnectionToDB(dbc *Config) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Port, dbc.DbName)
	db.Connection, _ = sql.Open("mysql", connectionString)
	db.Err = db.Connection.Ping()
	if db.Err != nil {
		log.Printf("error connection to database: \n%s\n", db.Err)
		db.Status = false
		return
	}
	log.Printf("Success connect to '%s' database\n", dbc.DbName)
	db.Status = true
	db.Err = nil
}

func (db *Db) InitDB() bool {
	if db.Status == false && db.Err != nil {
		log.Printf("Err, not connected to DB server")
		return false
	}
	tx, err := db.Connection.Begin()

	if err != nil {
		log.Printf("InitDB: err in begin transaction\n%s\n", err)
		return false
	}

	for _, query := range queriesForInitDb {
		_, err := tx.Exec(query)
		if err != nil {
			tx.Rollback()
			log.Printf("InitDB: err in transaction\n%s\n", err)
			return false
		}
	}

	if tx.Commit() != nil {
		log.Printf("InitDB: err in commit transaction\n%s\n", err)
		return false
	}
	log.Printf("Success init DB\n")
	return true
}


func (db *Db) FillDBTestData() bool{
	if db.Status == false && db.Err != nil {
		log.Printf("Err, not connected to DB server")
		return false
	}

	file, err := ioutil.ReadFile("./testData.sql")
	if err != nil {
		log.Printf("Err in script file: \n%s\n", err)
		return false
	}

	queries := strings.SplitAfter(string(file), ";")

	tx, err := db.Connection.Begin()

	if err != nil {
		log.Printf("FillDbTestData: err in begin transaction\n%s\n", err)
		return false
	}

	for _, query := range queries[:len(queries)-1] {
		_, err := tx.Exec(query)
		if err != nil {
			tx.Rollback()
			log.Printf("FillDbTestData: err in transaction\n%s\n", err)
			return false
		}
	}

	if tx.Commit() != nil {
		log.Printf("FillDbTestData: err in commit transaction\n%s\n", err)
		return false
	}
	log.Printf("Success fill test data in test db\n")
	return true
}

func (db *Db) DropDb(config *Config) {
	if db.Status == false && db.Err != nil {
		log.Printf("Err, not connected to DB")
		return
	}
	query := fmt.Sprintf("DROP DATABASE %s", config.DbName)
	_, err := db.Connection.Exec(query)
	if err != nil{
		log.Printf("Err in drop database '%s'\n%s\n ", config.DbName, err)
		return
	}
	log.Printf("Success drop table '%s'\n", config.DbName)
}
