package user

import ()

type Location struct {
	BotId   string
	GroupID string
	StateID string
}

type User struct {
	id         string
	ExternalId string
	state      Location
}

type StructCreateUser struct {
	externalId string
}

func CreateUser(input StructCreateUser) *User {
	user := User{ExternalId: input.externalId}
	dao := NewUserDAO(&user)
	dao.Save()
	return toModel(dao)
}
