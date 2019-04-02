package main

import (
	"./db"
	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var tdbc db.TestDBConfig
var tdb db.TestDb

func index(w http.ResponseWriter, r *http.Request) {

}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var p models.People
	var st models.Student
	var ac models.Accession
	var sd models.SensitiveData

	p.Fio = "Student"
	p.Password = "clearPass"
	p.Email = "st@gmail.com"
	p.PhoneNumber = "+0000000000"
	p.Sex = 1
	p.Comment = ""
	p.Status = "student"
	p.HaveAccess = true
	p.Birthday = "2006-01-01"

	st.DateAdmission = "2018-09-01"
	st.Semester = 1
	st.IsFullTime = true
	st.IsCut = false
	st.IdGroup = 2

	ac.GetAbsence = true

	sd.MilitaryId = "12kfsak"
	sd.RegAddress = "sajfjahsfja asf asjf kasjflas jfla sjfka"
	sd.Rntrs = "fks227176842190"
	sd.PassportCode = "we886589as"

}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	var p models.People
	var empl models.Employee
	var ac models.Accession
	var sd models.SensitiveData

	p.Fio = "Employee"
	p.Password = "clearPass"
	p.Email = "st@gmail.com"
	p.PhoneNumber = "+0000000000"
	p.Sex = 1
	p.Comment = ""
	p.Status = "student"
	p.HaveAccess = true
	p.Birthday = "1986-01-01"

	empl.IdCathedra = 1
	empl.IdRank = 2
	empl.DateInvite = "2006-09-01"

	ac.GetAbsence = true

	sd.MilitaryId = "12kfsak"
	sd.RegAddress = "sajfjahsfja asf asjf kasjflas jfla sjfka"
	sd.Rntrs = "fks227176842190"
	sd.PassportCode = "we886589as"

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/addStudent", addStudent).Methods("GET")
	r.HandleFunc("/addEmployee", addEmployee).Methods("GET")

	defer tdb.Connection.Close()

	tdbc.TestDBConfigInit("/testDBConfig.json")
	tdb.ConnectionToMysqlServer(&tdbc)
	tdb.CreateDB(tdbc.DbName)
	tdb.ConnectionToDB(&tdbc)
	tdb.InitDB(tdbc.SqlFile)
	tdb.FillDBTestData(tdbc.TestData)
	defer tdb.DropTestDB(tdbc.DbName)

	log.Fatal(http.ListenAndServe(":8080", r))
}
