package models

type AuthenStudent struct {
	Date	string 		`json:"date"`
	StateAuthen		bool	`json:"stateAuthen"`
}

type AuthenData struct {
	Id		int 		`json:"id"`
	NameStudent		string		`json:"nameStudent"`
	IdStudent		string		`json:"idStudent"`
	AuthenStudent   []AuthenStudent   `json:"authenStudent"`
}

type AuthenDatas []AuthenData