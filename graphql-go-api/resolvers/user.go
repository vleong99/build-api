package resolvers

import "graphql-go-api/models"

type User struct {
	userData models.User
}

func (u *User) Id() int32 {
	return u.userData.Id
}

func (u *User) Name() string {
	return u.userData.Name
}

func (u *User) Username() string {
	return u.userData.Username
}

func (u *User) Email() string {
	return u.userData.Email
}

func (u *User) Address() *Address {
	return &Address{addressData: u.userData.Address}
}
