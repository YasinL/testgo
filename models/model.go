package models

type User struct {
	Id       string `orm:"column(id);pk"`
	Name string
}


