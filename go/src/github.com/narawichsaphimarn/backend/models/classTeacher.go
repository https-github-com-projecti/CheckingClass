package models

type TClass struct {
	Id                int    `json:"id"`
	TClassName        string `json:"t_class_name"`
	TClassDescription string `json:"t_class_description"`
	TClassId          string `json:"t_class_id"`
	TUser             string `json:"user"`
}

type tClasss []TClass
