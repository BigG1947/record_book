package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int `json:"port"`
	DbName   string `json:"db_name"`
	SqlFile  string `json:"sql_file"`
}

type Db struct {
	Connection *sql.DB
	Status     bool
	Err        error
}

func (dbc *DBConfig) DBConfigInit(config string){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("Err check current dir: \n%s\n", err)
		return
	}

	file, err := ioutil.ReadFile(dir + config)
	if err != nil {
		log.Printf("Err in read config file: \n%s\n", err)
		return
	}

	err = json.Unmarshal(file, dbc)

	if err != nil {
		log.Printf("Err in decode json config file: \n%s\n", err)
		return
	}

	fmt.Printf("DBConfig successfull init!\n")
}

func (db *Db) ConnectionToMysqlServer(dbc *DBConfig) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%d)/", dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Port)
	db.Connection, _ = sql.Open("mysql", connectionString)
	db.Err = db.Connection.Ping()
	if db.Err != nil {
		log.Printf("error connection to mysql server: \n%s\n", db.Err)
		panic(db.Err.Error())
		db.Status = false
		return
	}
	db.Status = true
	db.Err = nil
}

func (db *Db) CreateDB(name string) bool {
	if db.Status == false && db.Err != nil {
		log.Printf("Err, not connected to MySQL server")
		return false
	}
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	_, err := db.Connection.Exec(query)
	if err != nil {
		log.Printf("error creating database \"%s\": \n%s\n", name, err)
		return false
	}
	return true
}

func (db *Db) ConnectionToDB(dbc *DBConfig) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Port, dbc.DbName)
	db.Connection, _ = sql.Open("mysql", connectionString)
	db.Err = db.Connection.Ping()
	if db.Err != nil {
		log.Printf("error connection to database: \n%s\n", db.Err)
		panic(db.Err.Error())
		db.Status = false
		return
	}
	db.Status = true
	db.Err = nil
}

func (db *Db) InitDB(script string) bool {
	if db.Status == false && db.Err != nil {
		log.Printf("Err, not connected to DB server")
		return false
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("Err check current dir: \n%s\n", err)
		return false
	}

	file, err := ioutil.ReadFile(dir + script)
	if err != nil {
		log.Printf("Err in script file: \n%s\n", err)
		return false
	}

	queries := strings.SplitAfter(string(file), ";")

	tx, err := db.Connection.Begin()

	if err != nil {
		log.Printf("InitDB: err in begin transaction\n%s\n", err)
		return false
	}

	for _, query := range queries[:len(queries)-1] {
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
	return true
}
