package models

type Accession struct {
	IdPeople     int  `json:"id_people"`
	EditAccess   bool `json:"edit_access"`
	SetAbsence   bool `json:"set_absence"`
	GetAbsence   bool `json:"get_absence"`
	SetMark      bool `json:"set_mark"`
	SetEvent     bool `json:"set_event"`
	GetSensitive bool `json:"get_sensitive"`
	SetSensitive bool `json:"set_sensitive"`
	GetYlist     bool `json:"get_ylist"`
	ManageAcadem bool `json:"manage_academ"`
}
