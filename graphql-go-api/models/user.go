package models

type User struct {
	Id       int32
	Name     string
	Username string
	Email    string
	Address  Address
}
