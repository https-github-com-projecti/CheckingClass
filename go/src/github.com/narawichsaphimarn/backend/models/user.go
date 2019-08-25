package models

type User struct {
	//ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TName      string             `json:"tName"`
	UserName   string             `json:"userName"`
	TId        string             `json:"tId"`
	TEmail     string             `json:"tEmail"`
	TWorkPlace string             `json:"tWorkPlace"`
	TPassword  string             `json:"tPassword"`
}

type Users []User
