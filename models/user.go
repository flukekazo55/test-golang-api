package models

type User struct {
	UserId   string `json:"userId" bson:"userId"`
	UserName string `json:"userName" bson:"userName"`
	Name     string `json:"name" bson:"name"`
}
