package models

type Qrcode struct {
	Time 	string `json:"time"`
	User 	string `json:"user"`
	Pass 	int    `json:"passOfCouse"`
}

type CreateQr struct {
	Id		int		`json:"id"`
	Qrcode	string	`json:"qrcode`
	Time 	string 	`json:"time"`
	Pass 	int    	`json:"passOfCouse"`
	TimeAuthen 	int 	`json:"time_authen"`
}

type newQr []Qrcode
type newCreate []CreateQr
