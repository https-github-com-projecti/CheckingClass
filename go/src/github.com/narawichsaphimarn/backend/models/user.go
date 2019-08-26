package models

type User struct {
	ID         int			  `json:"id"`
	TName      string             `json:"tName"`
	UserName   string             `json:"userName"`
	TId        string             `json:"tId"`
	TEmail     string             `json:"tEmail"`
	TWorkPlace string             `json:"tWorkPlace"`
	TPassword  string             `json:"tPassword"`
	TPicture   string             `json:"tPicture"`
}

type Users []User
