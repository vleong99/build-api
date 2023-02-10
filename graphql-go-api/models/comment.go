package models

type Comment struct {
	PostId int32
	ID     int32
	Name   string
	Email  string
	Body   string
}
