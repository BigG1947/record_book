package main

import (
	"time"
)

type UserData struct {
	//Password string
	ID           int             `json:"id"`
	Name         string          `json:"name"` // фио пользователя
	Img          string          `json:"img"`
	Sex          string          `json:"sex"`
	Birthday     string          `json:"birthday"`
	Blocked      bool            `json:"blocked"`
	Status       string          `json:"status"`
	Access       map[string]bool `json:"access"`
	IsStudent    bool            `json:"is_student"`
	StudentData  Student         `json:"student_data"`
	EmployeeData Employee        `json:"employee_data"`
	TTL          time.Time       `json:"ttl"`
}

type SensitiveData struct {
	Passport     string `json:"passport"`
	Registration string `json:"registration"`
	RNTRS        string `json:"rntrs"`
	MilitaryID   string `json:"military_id"`
}

type Student struct {
	DateJoin   string `json:"date_join"`
	IsFullTime bool   `json:"is_full_time"`
	IsCut      bool   `json:"is_cut"`
	Group      string `json:"group"`
	GroupID    int    `json:"group_id"`
	Semester   string `json:"semester"`
}

type Employee struct {
	DateJoinWork string `json:"date_join_work"`
	Position     string `json:"position"`
	OverseeGroup string `json:"oversee_group"`
	Department   string `json:"department"`
}

type User struct {
	Email  string
	Pass   string
	Access map[string]bool
}
