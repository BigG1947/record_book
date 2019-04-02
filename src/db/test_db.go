package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TestDBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
	SqlFile  string `json:"sql_file"`
	TestData string `json:"test_data"`
}

type TestDb struct {
	Connection *sql.DB
	Status     bool
	Err        error
}

func (tdbc *TestDBConfig) TestDBConfigInit(config string){
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

	err = json.Unmarshal(file, tdbc)

	if err != nil {
		log.Printf("Err in decode json config file: \n%s\n", err)
		return
	}

	fmt.Printf("TestDBConfig successfull init!\n")
}

func (tdb *TestDb) ConnectionToMysqlServer(tdbc *TestDBConfig) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%d)/", tdbc.User, tdbc.Password, tdbc.Protocol, tdbc.Host, tdbc.Port)
	tdb.Connection, _ = sql.Open("mysql", connectionString)
	tdb.Err = tdb .Connection.Ping()
	if tdb.Err != nil {
		log.Printf("error connection to mysql server: \n%s\n", tdb.Err)
		panic(tdb .Err.Error())
		tdb.Status = false
		return
	}
	tdb.Status = true
	tdb.Err = nil
	fmt.Printf("Successfull connection to MySQL server!\n")
}

func (tdb *TestDb) CreateDB(name string) bool {
	if tdb.Status == false && tdb.Err != nil {
		log.Printf("Err, not connected to MySQL server")
		return false
	}
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	_, err := tdb.Connection.Exec(query)
	if err != nil {
		log.Printf("error creating database \"%s\": \n%s\n", name, err)
		return false
	}
	fmt.Printf("Success create test db: %s\n", name)
	return true
}

func (tdb *TestDb) ConnectionToDB(tdbc *TestDBConfig) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", tdbc.User, tdbc.Password, tdbc.Protocol, tdbc.Host, tdbc.Port, tdbc.DbName)
	tdb.Connection, _ = sql.Open("mysql", connectionString)
	tdb.Err = tdb.Connection.Ping()
	if tdb.Err != nil {
		log.Printf("error connection to database: \n%s\n", tdb.Err)
		panic(tdb.Err.Error())
		tdb.Status = false
		return
	}
	tdb.Status = true
	tdb.Err = nil
	fmt.Printf("Success connection to '%s' db\n", tdbc.DbName)
}

func (tdb *TestDb) InitDB(script string) bool {
	if tdb.Status == false && tdb.Err != nil {
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

	tx, err := tdb.Connection.Begin()

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
	fmt.Printf("Success init db\n")
	return true
}

func (tdb *TestDb) FillDBTestData(script string) bool{
	if tdb.Status == false && tdb.Err != nil {
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

	tx, err := tdb.Connection.Begin()

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
	fmt.Printf("Success fill test data in test db\n")
	return true
}

func (tdb *TestDb) DropTestDB(name string)  {
	if tdb.Status == false && tdb.Err != nil {
		log.Printf("Err, not connected to DB server")
		return
	}
	query := fmt.Sprintf("DROP DATABASE %s;", name)
	_, err := tdb.Connection.Exec(query)
	if err != nil {
		log.Printf("error creating database \"%s\": \n%s\n", name, err)
		return
	}
	fmt.Printf("Success drop test db: %s\n", name)
	return
}
