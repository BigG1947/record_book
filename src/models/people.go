package models

type People struct {
	Id          int    `json:"id"`
	Fio         string `json:"fio"`
	Birthday    string `json:"birthday"`
	Sex         int    `json:"sex"`
	Comment     string `json:"comment"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	HaveAccess  bool   `json:"have_access"`
}