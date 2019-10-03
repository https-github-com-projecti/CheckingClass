package models

type TClass struct {
	Id                int    `json:"subject_id"`
	TClassName        string `json:"TSName"`
	TClassDescription string `json:"TSDescription"`
	TClassId          string `json:"TSID"`
	TUser             string `json:"TSTeacher"`
	TClassPass 		  int    `json:"TSpassword"`
	
}

type tClasss []TClass
