package models

type Employee struct {
	IdPeople   int    `json:"id_people"`
	DateInvite string `json:"date_invite"`
	IdRank     int    `json:"id_rank"`
	IdGroup    int    `json:"id_group"`
	IdCathedra int    `json:"id_cathedra"`
}