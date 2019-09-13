package models

type User struct {
	ID         		int			       `json:"id"`
	TFirstName      string        	   `json:"tFirstName"`
	TLastName		string		 	   `json:"tLastName"`
	UserName   		string             `json:"userName"`
	TId        		string             `json:"tId"`
	TEmail     		string             `json:"tEmail"`
	TWorkPlace 		string             `json:"tWorkPlace"`
	TPassword  		string             `json:"tPassword"`
	TPicture   		string             `json:"tPicture"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users []User
type UserLogins []UserLogin
