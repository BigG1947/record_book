package main

import (
	dbPack "./db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var db dbPack.Db

func index(w http.ResponseWriter, r *http.Request) {

}

func getStudent(w http.ResponseWriter, r *http.Request) {
	students, err := dbPack.GetAllStudent(&db)

	if err != nil {
		fmt.Fprintf(w, "Error: \n%s\n", err)
		return
	}

	for _, student := range students {
		var g dbPack.Group

		err := g.GetGroupById(db.Connection, student.Student.IdGroup)
		if err != nil{fmt.Printf("Err: %s\n", err.Error())}

		fmt.Fprintf(w, "--------- %s --------\nid: %d\nfio: %s\nbirthday: %s\ngender: %d\nimg: %s\ncomment: %s\npassword: %s\nphone_number: %s\nemail: %s\nstatus: %s\nhave_access: %v\n\n", student.Fio, student.Id, student.Fio, student.Birthday, student.Gender, student.Img, student.Comment, student.Password, student.PhoneNumber, student.Email, student.Status, student.HaveAccess)
		fmt.Fprintf(w, "Student Info:\nid_people: %d\ndate_addmission: %s\nis_full_time: %v\nis_cut: %v\nid_group: %d\nsemester: %d\n\n", student.Student.IdPeople, student.Student.DateAdmission, student.Student.IsFullTime, student.Student.IsCut, student.Student.IdGroup, student.Student.Semester)
		fmt.Fprintf(w, "Group Info:\nid_group: %d\nname: %s\nid_direction: %d\n\n", g.Id, g.Name, g.IdDirection)
		fmt.Fprintf(w, "Accession:\nid_people: %d\nedit_access: %v\nset_absence: %v\nget_absence: %v\nset_mark: %v\nset_envent: %v\nget_sesnsitive: %v\nset_sensitive: %v\nget_ylist: %v\nmanage_academ: %v\n\n", student.Accession.IdPeople, student.Accession.EditAccess, student.Accession.SetAbsence, student.Accession.GetAbsence, student.Accession.SetMark, student.Accession.SetEvent, student.Accession.GetSensitive, student.Accession.SetSensitive, student.Accession.GetYlist, student.Accession.ManageAcadem)
		fmt.Fprintf(w, "Sensitive Data:\nid_people: %d\npassport_code: %s\nrntrs: %s\nreg_address: %s\nmillitary_id: %s\n\n", student.SensitiveData.IdPeople, student.SensitiveData.PassportCode, student.SensitiveData.Rntrs, student.SensitiveData.RegAddress, student.SensitiveData.MilitaryId)
	}
	return
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	employees, err := dbPack.GetAllEmployee(&db)
	if err != nil {
		fmt.Fprintf(w, "Error: \n%s\n", err)
		return
	}

	for _, employee := range employees {
		var g dbPack.Group
		if employee.Employee.IdGroup != 0{
			err := g.GetGroupById(db.Connection, employee.Employee.IdGroup)
			if err != nil{fmt.Printf("Err: %s\n", err.Error())}
		}
		fmt.Fprintf(w, "--------- %s --------\nid: %d\nfio: %s\nbirthday: %s\ngender: %d\nimg: %s\ncomment: %s\npassword: %s\nphone_number: %s\nemail: %s\nstatus: %s\nhave_access: %v\n\n", employee.Fio, employee.Id, employee.Fio, employee.Birthday, employee.Gender, employee.Img, employee.Comment, employee.Password, employee.PhoneNumber, employee.Email, employee.Status, employee.HaveAccess)
		fmt.Fprintf(w, "Employee Info:\nid_people: %d\ndate_invite: %s\nid_rank: %d\nid_group: %d\nid_cathedra: %d\n\n", employee.Employee.IdPeople, employee.Employee.DateInvite, employee.Employee.IdRank, employee.Employee.IdGroup, employee.Employee.IdCathedra)
		fmt.Fprintf(w, "Group Info:\nid_group: %d\nname: %s\nid_direction: %d\n\n", g.Id, g.Name, g.IdDirection)
		fmt.Fprintf(w, "Accession:\nid_people: %d\nedit_access: %v\nset_absence: %v\nget_absence: %v\nset_mark: %v\nset_envent: %v\nget_sesnsitive: %v\nset_sensitive: %v\nget_ylist: %v\nmanage_academ: %v\n\n", employee.Accession.IdPeople, employee.Accession.EditAccess, employee.Accession.SetAbsence, employee.Accession.GetAbsence, employee.Accession.SetMark, employee.Accession.SetEvent, employee.Accession.GetSensitive, employee.Accession.SetSensitive, employee.Accession.GetYlist, employee.Accession.ManageAcadem)
		fmt.Fprintf(w, "Sensitive Data:\nid_people: %d\npassport_code: %s\nrntrs: %s\nreg_address: %s\nmillitary_id: %s\n\n", employee.SensitiveData.IdPeople, employee.SensitiveData.PassportCode, employee.SensitiveData.Rntrs, employee.SensitiveData.RegAddress, employee.SensitiveData.MilitaryId)
	}
	return
}

func getSinglePeople(w http.ResponseWriter, r *http.Request) {
	var p dbPack.People

	p.GetPeopleByEmail(&db, "ivanov@gmail.com")

	fmt.Fprintf(w, "--------- %s --------\nid: %d\nfio: %s\nbirthday: %s\ngender: %d\nimg: %s\ncomment: %s\npassword: %s\nphone_number: %s\nemail: %s\nstatus: %s\nhave_access: %v\n\n", p.Fio, p.Id, p.Fio, p.Birthday, p.Img, p.Gender, p.Comment, p.Password, p.PhoneNumber, p.Email, p.Status, p.HaveAccess)
	fmt.Fprintf(w, "Student Info:\nid_people: %d\ndate_addmission: %s\nis_full_time: %v\nis_cut: %v\nid_group: %d\nsemester: %d\n\n", p.Student.IdPeople, p.Student.DateAdmission, p.Student.IsFullTime, p.Student.IsCut, p.Student.IdGroup, p.Student.Semester)
	fmt.Fprintf(w, "Employee Info:\nid_people: %d\ndate_invite: %s\nid_rank: %d\nid_group: %d\nid_cathedra: %d\n\n", p.Employee.IdPeople, p.Employee.DateInvite, p.Employee.IdRank, p.Employee.IdGroup, p.Employee.IdCathedra)
	fmt.Fprintf(w, "Accession:\nid_people: %d\nedit_access: %v\nset_absence: %v\nget_absence: %v\nset_mark: %v\nset_envent: %v\nget_sesnsitive: %v\nset_sensitive: %v\nget_ylist: %v\nmanage_academ: %v\n\n", p.Accession.IdPeople, p.Accession.EditAccess, p.Accession.SetAbsence, p.Accession.GetAbsence, p.Accession.SetMark, p.Accession.SetEvent, p.Accession.GetSensitive, p.Accession.SetSensitive, p.Accession.GetYlist, p.Accession.ManageAcadem)
	fmt.Fprintf(w, "Sensitive Data:\nid_people: %d\npassport_code: %s\nrntrs: %s\nreg_address: %s\nmillitary_id: %s\n\n", p.SensitiveData.IdPeople, p.SensitiveData.PassportCode, p.SensitiveData.Rntrs, p.SensitiveData.RegAddress, p.SensitiveData.MilitaryId)

}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var p dbPack.People
	var st dbPack.Student
	var ac dbPack.Accession
	var sd dbPack.SensitiveData

	p.Fio = "Student"
	p.Password = "clearPass"
	p.Email = "st@gmail.com"
	p.PhoneNumber = "+0000000000"
	p.Gender = 1
	p.Img = "/static/img/default.png"
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

	p.SensitiveData = sd
	p.Student = st
	p.Accession = ac

	p.InsertStudent(&db)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	var p dbPack.People
	var empl dbPack.Employee
	var ac dbPack.Accession
	var sd dbPack.SensitiveData

	p.Fio = "Employee"
	p.Password = "clearPass"
	p.Email = r.URL.Query()["email"][0]
	p.PhoneNumber = "+0000000000"
	p.Gender = 1
	p.Img = "/static/img/default.png"
	p.Comment = ""
	p.Status = "employee"
	p.HaveAccess = true
	p.Birthday = "1986-01-01"

	empl.IdCathedra = 1
	empl.IdRank = 2
	empl.IdGroup, _ = strconv.Atoi(r.URL.Query().Get("id_group"))
	empl.DateInvite = "2006-09-01"

	ac.GetAbsence = true

	sd.MilitaryId = "12kfsak"
	sd.RegAddress = "sajfjahsfja asf asjf kasjflas jfla sjfka"
	sd.Rntrs = "fks227176842190"
	sd.PassportCode = "we886589as"

	p.SensitiveData = sd
	p.Employee = empl
	p.Accession = ac

	p.InsertEmployee(&db)

}

func login(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("success")
	fmt.Fprintf(w, string(res))
}

func printMap(w http.ResponseWriter, r *http.Request) {
	authInfo := dbPack.GetEmailPasswordMap(&db)

	fmt.Fprintf(w, "Email: Password\n")
	for key, value := range authInfo {
		fmt.Fprintf(w, "%s: %s\n", key, value)
	}
}

func angularJs(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
	return
}

func getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := dbPack.GetAllGroup(db.Connection)
	if err != nil{
		fmt.Fprintf(w, "Err: %s\n", err.Error())
		return
	}
	for _, group := range groups{
		fmt.Fprintf(w, "--------- %s --------\nid: %d\nname: %s\nid_direction: %d\n\n", group.Name, group.Id, group.Name, group.IdDirection)
	}
	return
}

func getRanks(w http.ResponseWriter, r *http.Request)  {
	ranks, err := dbPack.GetAllRanks(db.Connection)
	if err != nil{
		fmt.Fprintf(w, "Err: %s\n", err.Error())
		return
	}
	for _, rank := range ranks{
		fmt.Fprintf(w, "--------- %s --------\nid: %d\nname: %s\n\n", rank.Name, rank.Id, rank.Name)
	}
	return
}



func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/angular", angularJs)

	r.HandleFunc("/", index)

	r.HandleFunc("/addStudent", addStudent).Methods("GET")
	r.HandleFunc("/getStudent", getStudent).Methods("GET")

	r.HandleFunc("/addEmployee", addEmployee).Methods("GET")
	r.HandleFunc("/getEmployee", getEmployee).Methods("GET")

	r.HandleFunc("/getGroups", getGroups).Methods("GET")

	r.HandleFunc("/getRanks", getRanks).Methods("GET")


	r.HandleFunc("/single", getSinglePeople).Methods("GET")
	r.HandleFunc("/printMap", printMap).Methods("GET")

	db.ConnectionToMysqlServer(&dbPack.DefaultConfig)
	db.DropDb(&dbPack.DefaultConfig)
	db.CreateDB(dbPack.DefaultConfig.DbName)
	db.ConnectionToDB(&dbPack.DefaultConfig)
	db.InitDB()
	db.FillDBTestData()

	log.Fatal(http.ListenAndServe(":8080", r))
}