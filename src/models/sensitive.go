package models

type SensitiveData struct {
	IdPeople     int    `json:"id_people"`
	PassportCode string `json:"passport_code"`
	Rntrs        string `json:"rntrs"`
	RegAddress   string `json:"reg_address"`
	MilitaryId   string `json:"military_id"`
}