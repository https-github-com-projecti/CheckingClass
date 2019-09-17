package models

type Qrcode struct {
	Time string `json:"time"`
	User string `json:"user"`
	Pass string    `json:"passOfCouse"`
}

type CreateQr struct {
	Id		int		`json:"id"`
	Qrcode	string	`json:"qrcode`
	Time 	string 	`json:"time"`
	User 	string 	`json:"user"`
}

type newQr []Qrcode
type newCreate []CreateQr
