package models

type Qrcode struct {
	Time string `json:"time"`
	User string `json:"user"`
	Pass int    `json:"passOfCouse"`
}

type CreateQr struct {
	Id         int    `json:"id"`
	Qrcode     string `json:"AQRcode"`
	Time       string `json:"ADate"`
	Pass       int    `json:"passOfCouse"`
	TimeAuthen int    `json:"ATimeAuthen"`
}

type newQr []Qrcode
type newCreate []CreateQr
