package main

import (
	"./blockchain"
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
var bc *blockchain.BlockChain

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/feedbackForm.html")
	if err != nil {
		fmt.Fprintf(w, "Err in templating: %s", err)
		return
	}
	empl, err := dbPack.GetAllEmployee(db.Connection)
	if err != nil {
		fmt.Fprintf(w, "Err: %s", err)
		return
	}
	tmpl.Execute(w, empl)
	return
}

func printBC(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	bci := bc.Iterator()
	bc.PrintBlockChain(bci, &w)
	return
}

func getFeedback(w http.ResponseWriter, r *http.Request) {
	feedbacks, err := dbPack.GetAllFeedBacks(db.Connection)
	if err != nil {
		fmt.Fprintf(w, "Error при запросе отзывов: %s\n", err)
		return
	}
	for _, f := range feedbacks {
		var p dbPack.People
		p.GetPeopleById(db.Connection, f.EmployeeId)
		fmt.Fprintf(w, "Employee: %s\nFeedback: %s\nMark: %d\n\n", p.Fio, f.Data, f.Mark)
	}
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	students, err := dbPack.GetAllStudent(db.Connection)

	if err != nil {
		fmt.Fprintf(w, "Error: \n%s\n", err)
		return
	}

	for _, student := range students {
		fmt.Fprintf(w, "--------- %s --------\nid: %d\nfio: %s\nbirthday: %s\ngender: %d\nimg: %s\ncomment: %s\npassword: %s\nphone_number: %s\nemail: %s\nstatus: %s\nhave_access: %v\n\n", student.Fio, student.Id, student.Fio, student.Birthday, student.Gender, student.Img, student.Comment, student.Password, student.PhoneNumber, student.Email, student.Status, student.HaveAccess)
		fmt.Fprintf(w, "Status info:\nid: %d\nname:%s\n\n", student.Status.Id, student.Status.Name)
		fmt.Fprintf(w, "Student Info:\nid_people: %d\ndate_addmission: %s\nis_full_time: %v\nis_cut: %v\nid_group: %d\nsemester: %d\n\n", student.Student.IdPeople, student.Student.DateAdmission, student.Student.IsFullTime, student.Student.IsCut, student.Student.Group.Id, student.Student.Semester)
		fmt.Fprintf(w, "Group Info:\nid_group: %d\nname: %s\nid_direction: %d\n\n", student.Student.Group.Id, student.Student.Group.Name, student.Student.Group.IdDirection)
		fmt.Fprintf(w, "Student Mark:\n")
		for _, mark := range student.Student.Marks {
			fmt.Fprintf(w, "id_disciplie: %d\nid_employee: %d\nvalue: %d\nnational_value: %s\nis_exam: %v\nsemester: %d\ndate: %s\n\n", mark.IdDiscipline, mark.IdEmployee, mark.Value, mark.NationalValue, mark.IsExam, mark.Semester, mark.Date)
		}
		fmt.Fprintf(w, "Accession:\nid_people: %d\nedit_access: %v\nset_absence: %v\nget_absence: %v\nset_mark: %v\nset_envent: %v\nget_sesnsitive: %v\nset_sensitive: %v\nget_ylist: %v\nmanage_academ: %v\n\n", student.Accession.IdPeople, student.Accession.EditAccess, student.Accession.SetAbsence, student.Accession.GetAbsence, student.Accession.SetMark, student.Accession.SetEvent, student.Accession.GetSensitive, student.Accession.SetSensitive, student.Accession.GetYlist, student.Accession.ManageAcadem)
		//fmt.Fprintf(w, "Sensitive Data:\nid_people: %d\npassport_code: %s\nrntrs: %s\nreg_address: %s\nmillitary_id: %s\n\n", student.SensitiveData.IdPeople, student.SensitiveData.PassportCode, student.SensitiveData.Rntrs, student.SensitiveData.RegAddress, student.SensitiveData.MilitaryId)
	}
	return
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	employees, err := dbPack.GetAllEmployee(db.Connection)
	if err != nil {
		fmt.Fprintf(w, "Error: \n%s\n", err)
		return
	}

	for _, employee := range employees {
		fmt.Fprintf(w, "--------- %s --------\nid: %d\nfio: %s\nbirthday: %s\ngender: %d\nimg: %s\ncomment: %s\npassword: %s\nphone_number: %s\nemail: %s\nstatus: %s\nhave_access: %v\n\n", employee.Fio, employee.Id, employee.Fio, employee.Birthday, employee.Gender, employee.Img, employee.Comment, employee.Password, employee.PhoneNumber, employee.Email, employee.Status, employee.HaveAccess)
		fmt.Fprintf(w, "Status info:\nid: %d\nname:%s\n\n", employee.Status.Id, employee.Status.Name)
		fmt.Fprintf(w, "Employee Info:\nid_people: %d\ndate_invite: %s\nid_rank: %d\nid_group: %d\nid_cathedra: %d\n\n", employee.Employee.IdPeople, employee.Employee.DateInvite, employee.Employee.Rank.Id, len(employee.Employee.Group), employee.Employee.Cathedra.Id)
		fmt.Fprintf(w, "Group Info:\n")
		for _, g := range employee.Employee.Group {
			fmt.Fprintf(w, "id_group: %d\nname: %s\nid_direction: %d\n\n", g.Id, g.Name, g.IdDirection)
		}
		fmt.Fprintf(w, "Employee Discipline:\n")
		for _, d := range employee.Employee.Disciplines {
			fmt.Fprintf(w, "name: %s\n\n", d.Name)
		}
		fmt.Fprintf(w, "Rank info:\nid: %d\nname:%s\n\n", employee.Employee.Rank.Id, employee.Employee.Rank.Name)
		fmt.Fprintf(w, "Cathedra info:\nid: %d\nname:%s\n\n", employee.Employee.Cathedra.Id, employee.Employee.Cathedra.Name)
		fmt.Fprintf(w, "Accession:\nid_people: %d\nedit_access: %v\nset_absence: %v\nget_absence: %v\nset_mark: %v\nset_envent: %v\nget_sesnsitive: %v\nset_sensitive: %v\nget_ylist: %v\nmanage_academ: %v\n\n", employee.Accession.IdPeople, employee.Accession.EditAccess, employee.Accession.SetAbsence, employee.Accession.GetAbsence, employee.Accession.SetMark, employee.Accession.SetEvent, employee.Accession.GetSensitive, employee.Accession.SetSensitive, employee.Accession.GetYlist, employee.Accession.ManageAcadem)
		//fmt.Fprintf(w, "Sensitive Data:\nid_people: %d\npassport_code: %s\nrntrs: %s\nreg_address: %s\nmillitary_id: %s\n\n", employee.SensitiveData.IdPeople, employee.SensitiveData.PassportCode, employee.SensitiveData.Rntrs, employee.SensitiveData.RegAddress, employee.SensitiveData.MilitaryId)
	}
	return
}

func getSinglePeople(w http.ResponseWriter, r *http.Request) {
	var p dbPack.People

	p.GetPeopleByEmail(db.Connection, "admin@admin.com")

	fmt.Fprintf(w, "--------- %s --------\nid: %d\nfio: %s\nbirthday: %s\ngender: %d\nimg: %s\ncomment: %s\npassword: %s\nphone_number: %s\nemail: %s\nstatus: %s\nhave_access: %v\n\n", p.Fio, p.Id, p.Fio, p.Birthday, p.Img, p.Gender, p.Comment, p.Password, p.PhoneNumber, p.Email, p.Status, p.HaveAccess)
	fmt.Fprintf(w, "Student Info:\nid_people: %d\ndate_addmission: %s\nis_full_time: %v\nis_cut: %v\nid_group: %d\nsemester: %d\n\n", p.Student.IdPeople, p.Student.DateAdmission, p.Student.IsFullTime, p.Student.IsCut, p.Student.Group.Id, p.Student.Semester)
	fmt.Fprintf(w, "Employee Info:\nid_people: %d\ndate_invite: %s\nid_rank: %d\nid_group: %d\nid_cathedra: %d\n\n", p.Employee.IdPeople, p.Employee.DateInvite, p.Employee.Rank.Id, len(p.Employee.Group), p.Employee.Cathedra.Id)
	for _, g := range p.Employee.Group {
		fmt.Fprintf(w, "Group Info:\nid_group: %d\nname: %s\nid_direction: %d\n\n", g.Id, g.Name, g.IdDirection)
	}
	fmt.Fprintf(w, "Accession:\nid_people: %d\nedit_access: %v\nset_absence: %v\nget_absence: %v\nset_mark: %v\nset_envent: %v\nget_sesnsitive: %v\nset_sensitive: %v\nget_ylist: %v\nmanage_academ: %v\n\n", p.Accession.IdPeople, p.Accession.EditAccess, p.Accession.SetAbsence, p.Accession.GetAbsence, p.Accession.SetMark, p.Accession.SetEvent, p.Accession.GetSensitive, p.Accession.SetSensitive, p.Accession.GetYlist, p.Accession.ManageAcadem)
	//fmt.Fprintf(w, "Sensitive Data:\nid_people: %d\npassport_code: %s\nrntrs: %s\nreg_address: %s\nmillitary_id: %s\n\n", p.SensitiveData.IdPeople, p.SensitiveData.PassportCode, p.SensitiveData.Rntrs, p.SensitiveData.RegAddress, p.SensitiveData.MilitaryId)

}

func getAvg(w http.ResponseWriter, r *http.Request) {
	var avgfeedback []dbPack.AvgFeedback
	var err error
	idGroup, err := strconv.Atoi(r.URL.Query().Get("group"))
	if err != nil {
		avgfeedback, err = dbPack.GetAllFeedBackWithAVG(db.Connection)
	} else {
		avgfeedback, err = dbPack.GetAllFeedBackWithAVGByStudent(db.Connection, idGroup)
	}

	if err != nil {
		fmt.Fprintf(w, "err in avg: %s\n", err)
		return
	}
	for _, af := range avgfeedback {
		fmt.Fprintf(w, "id_employee: %d\nname: %s\navg_mark: %d\n\n", af.IdEmployee, af.NameEmployee, af.AvgMark)
	}
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var p dbPack.People
	var st dbPack.Student
	var ac dbPack.Accession
	var sd dbPack.SensitiveData

	p.Fio = "Student"
	p.Password = "clearPass"
	p.Email = "st531@gmail.com"
	p.PhoneNumber = "+0000000000"
	p.Gender = 1
	p.Img = "/static/img/default.png"
	p.Comment = ""
	p.Status.Id = 1
	p.HaveAccess = true
	p.Birthday = "2006-01-01"

	st.DateAdmission = "2018-09-01"
	st.Semester = 1
	st.IsFullTime = true
	st.IsCut = false
	st.Group.Id = 2

	ac.GetAbsence = true

	sd.MilitaryId = "12kfsak"
	sd.RegAddress = "sajfjahsfja asf asjf kasjflas jfla sjfka"
	sd.Rntrs = "fks227176842190"
	sd.PassportCode = "we886589as"

	//p.SensitiveData = sd
	p.Student = st
	p.Accession = ac

	err := p.InsertStudent(db.Connection)
	fmt.Fprint(w, "%v", err)
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
	p.Status.Id = 1
	p.HaveAccess = true
	p.Birthday = "1986-01-01"

	empl.Cathedra.Id = 1
	empl.Rank.Id = 2
	empl.DateInvite = "2006-09-01"

	ac.GetAbsence = true

	sd.MilitaryId = "12kfsak"
	sd.RegAddress = "sajfjahsfja asf asjf kasjflas jfla sjfka"
	sd.Rntrs = "fks227176842190"
	sd.PassportCode = "we886589as"

	//p.SensitiveData = sd
	p.Employee = empl
	p.Accession = ac

	err := p.InsertEmployee(db.Connection)
	if err != nil {
		fmt.Fprint(w, "Error: %v", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("success")
	fmt.Fprintf(w, string(res))
}

func printMap(w http.ResponseWriter, r *http.Request) {
	authInfo, err := dbPack.GetEmailPasswordMap(db.Connection)
	if err != nil {
		fmt.Fprintf(w, "Err: %s", err.Error())
		return
	}
	fmt.Fprintf(w, "Email: Password\n")
	for key, value := range authInfo {
		fmt.Fprintf(w, "%s: %s\n", key, value)
	}
}

func angularJs(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("/templates/index.html")
	tmpl.Execute(w, nil)
	return
}

func getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := dbPack.GetAllGroup(db.Connection)
	if err != nil {
		fmt.Fprintf(w, "Err: %s\n", err.Error())
		return
	}
	for _, group := range groups {
		fmt.Fprintf(w, "--------- %s --------\nid: %d\nname: %s\nid_direction: %d\n\n", group.Name, group.Id, group.Name, group.IdDirection)
	}
	return
}

func getRanks(w http.ResponseWriter, r *http.Request) {
	ranks, err := dbPack.GetAllRanks(db.Connection)
	if err != nil {
		fmt.Fprintf(w, "Err: %s\n", err.Error())
		return
	}
	for _, rank := range ranks {
		fmt.Fprintf(w, "--------- %s --------\nid: %d\nname: %s\n\n", rank.Name, rank.Id, rank.Name)
	}
	return
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	r.HandleFunc("/blockchain/checkActivity", checkActivity).Methods("GET")

	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/angular", angularJs)

	r.HandleFunc("/", index)
	r.HandleFunc("/addFeedBack", addFeedback).Methods("POST")
	r.HandleFunc("/getFeedBack", getFeedback)
	r.HandleFunc("/printBC", printBC)
	r.HandleFunc("/getAvg", getAvg)

	r.HandleFunc("/addStudent", addStudent).Methods("GET")
	r.HandleFunc("/getStudent", getStudent).Methods("GET")

	r.HandleFunc("/addEmployee", addEmployee).Methods("GET")
	r.HandleFunc("/getEmployee", getEmployee).Methods("GET")

	r.HandleFunc("/getGroups", getGroups).Methods("GET")

	r.HandleFunc("/getRanks", getRanks).Methods("GET")

	r.HandleFunc("/single", getSinglePeople).Methods("GET")
	r.HandleFunc("/printMap", printMap).Methods("GET")

	err := db.TestConnection(&dbPack.DefaultConfigLaptop)
	if err != nil {
		fmt.Printf("Error in test connection: %s\n", err)
	}

	bc, err = blockchain.InitBlockChain(db.Connection)
	if err != nil {
		fmt.Printf("Error in init BlockChain: %s\n", err)
	}
	bc.FillTestFeedBack()

	test()

	log.Printf("Starting server...\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func test() {
	var err error
	var institute dbPack.Institute
	institute.Name = "Новый институт"
	institute.Id, err = dbPack.InsertInstitute(db.Connection, &institute)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	institute.Name = "This Work!"
	err = dbPack.UpdateInstitute(db.Connection, &institute)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	err = dbPack.DeleteInstitute(db.Connection, institute.Id)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	var faculty dbPack.Faculty
	faculty.Name = "Новый факультет"
	faculty.IdInstitute = 1
	faculty.Id, err = dbPack.InsertFaculty(db.Connection, &faculty)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	faculty.Name = "This Work!"
	err = dbPack.UpdateFaculty(db.Connection, &faculty)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	err = dbPack.DeleteFaculty(db.Connection, faculty.Id)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	var cathedra dbPack.Cathedra
	cathedra.Name = "Новая кафедра"
	cathedra.IdFaculty = 1
	cathedra.Id, err = dbPack.InsertCathedra(db.Connection, &cathedra)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	cathedra.Name = "This Work!"
	err = dbPack.UpdateCathedra(db.Connection, &cathedra)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	err = dbPack.DeleteCathedra(db.Connection, cathedra.Id)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	var direction dbPack.Direction
	direction.Name = "Новое направление"
	direction.IdCathedra = 1
	direction.Id, err = dbPack.InsertDirection(db.Connection, &direction)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	direction.Name = "This Work!"
	err = dbPack.UpdateDirection(db.Connection, &direction)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	err = dbPack.DeleteDirection(db.Connection, direction.Id)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	var group dbPack.Group
	group.Name = "Новая группа"
	group.IdDirection = 1
	group.Id, err = group.Insert(db.Connection)
	if err != nil {
		fmt.Printf("Err group.Insert: %s\n", err)
		return
	}
	group.Name = "This Work!"
	group.IdDirection = 2
	err = dbPack.UpdateGroup(db.Connection, &group)
	if err != nil {
		fmt.Printf("Err UpdateGroup: %s\n", err)
		return
	}
	err = dbPack.DeleteGroup(db.Connection, group.Id)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	var rank dbPack.Rank
	rank.Name = "Новая должность"
	rank.Id, err = dbPack.InsertRank(db.Connection, &rank)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	rank.Name = "This Work!"
	err = dbPack.UpdateRank(db.Connection, &rank)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	err = dbPack.DeleteRank(db.Connection, rank.Id)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	_, err = dbPack.GetAllLoadsForEmployee(db.Connection, 3)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	_, err = dbPack.GetAllLoadsForAssistent(db.Connection, 6)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	_, err = dbPack.GetAllLoadsByIdGroup(db.Connection, 1)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	var l dbPack.Load
	l.IdEmployee = 3
	l.IdAssistant = 6
	l.IdGroup = 1
	l.Discipline.Id = 1
	l.Semester.Id = 1
	l.Id, err = l.Insert(db.Connection)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	l.Semester.Id = 2
	err = dbPack.UpdateLoadsById(db.Connection, &l)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	err = dbPack.DeleteLoadsById(db.Connection, l.Id)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	_, err = dbPack.GetAllDisciplineForEmployee(db.Connection, 3)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	_, err = dbPack.GetAllLoadsForEmployee(db.Connection, 3)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	_, err = dbPack.GetStudentDisciplinesByMark(db.Connection, 1)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	_, err = dbPack.GetCurrentStudentDiscipline(db.Connection, 1)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	_, err = dbPack.GetAllLoadsByIdGroup(db.Connection, 1)
	if err != nil {
		fmt.Printf("GetAllLoadsByIdGroup. Err: %s\n", err)
		return
	}
	_, err = dbPack.GetAllLoadsForEmployee(db.Connection, 1)
	if err != nil {
		fmt.Printf("GetAllLoadsForEmployee. Err: %s\n", err)
		return
	}
	_, err = dbPack.GetAllLoadsForAssistent(db.Connection, 1)
	if err != nil {
		fmt.Printf("GetAllLoadsForAssistent. Err: %s\n", err)
		return
	}
	_, err = dbPack.GetAllStudent(db.Connection)
	if err != nil {
		fmt.Printf("Err GetAllStudent: %s\n", err)
		return
	}
	_, err = dbPack.GetAllEmployee(db.Connection)
	if err != nil {
		fmt.Printf("Err GetAllEmployee: %s\n", err)
		return
	}

	fmt.Printf("Ok!\n")
	return
}
