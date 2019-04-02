package models

type Student struct {
	IdPeople      int    `json:"id_people"`
	DateAdmission string `json:"date_admission"`
	IsFullTime    bool   `json:"is_full_time"`
	IsCut         bool   `json:"is_cut"`
	IdGroup       int    `json:"id_group"`
	Semester      int    `json:"semester"`
}