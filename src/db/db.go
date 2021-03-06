package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
)

type Db struct {
	Connection *sql.DB
	Status     bool
	Err        error
}

func (db *Db) TestConnection(dbc *Config) error {
	db.ConnectionToMysqlServer(dbc)
	if db.Err != nil {
		return db.Err
	}
	_, err := db.DropDb(dbc)
	if err != nil {
		return err
	}
	_, err = db.CreateDB(dbc.DbName)
	if err != nil {
		return err
	}
	db.ConnectionToDB(dbc)
	if db.Err != nil {
		return db.Err
	}
	_, err = db.InitDB()
	if err != nil {
		return err
	}
	_, err = db.FillFirstData()
	if err != nil {
		return err
	}
	_, err = db.FillDBTestData()
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) ConnectionToMysqlServer(dbc *Config) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%d)/", dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Port)
	db.Connection, _ = sql.Open("mysql", connectionString)
	db.Err = db.Connection.Ping()
	if db.Err != nil {
		db.Status = false
		return
	}
	db.Status = true
	db.Err = nil
}

func (db *Db) CreateDB(name string) (bool, error) {
	if db.Status == false && db.Err != nil {
		return db.Status, db.Err
	}
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	_, err := db.Connection.Exec(query)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (db *Db) ConnectionToDB(dbc *Config) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Port, dbc.DbName)
	db.Connection, _ = sql.Open("mysql", connectionString)
	db.Err = db.Connection.Ping()
	if db.Err != nil {
		db.Status = false
		return
	}
	db.Status = true
	db.Err = nil
}

func (db *Db) InitDB() (bool, error) {
	if db.Status == false && db.Err != nil {
		return db.Status, db.Err
	}
	tx, err := db.Connection.Begin()

	if err != nil {
		return false, err
	}

	for _, query := range queriesForInitDb {
		_, err := tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	if tx.Commit() != nil {
		return false, err
	}
	return true, err
}

func (db *Db) FillFirstData() (bool, error) {
	if db.Status == false && db.Err != nil {
		return db.Status, db.Err
	}
	tx, err := db.Connection.Begin()

	if err != nil {
		return false, err
	}

	for _, query := range FillDataQueries {
		_, err := tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	if tx.Commit() != nil {
		return false, err
	}
	return true, err
}

func (db *Db) FillDBTestData() (bool, error) {
	if db.Status == false && db.Err != nil {
		return db.Status, db.Err
	}

	file, err := ioutil.ReadFile("./testData.sql")
	if err != nil {
		return false, err
	}

	queries := strings.SplitAfter(string(file), ";")

	tx, err := db.Connection.Begin()

	if err != nil {
		return false, err
	}

	for _, query := range queries[:len(queries)-1] {
		_, err := tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	if tx.Commit() != nil {
		return false, err
	}

	return true, err
}

func (db *Db) DropDb(config *Config) (bool, error) {
	if db.Status == false && db.Err != nil {
		return false, db.Err
	}
	query := fmt.Sprintf("DROP DATABASE %s", config.DbName)
	_, err := db.Connection.Exec(query)
	if err != nil {
		return false, err
	}
	return true, nil
}
