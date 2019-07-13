package user

func NewUser() user {
	return user{}
}

type user struct {
	level TypeUser
}

func (el *user) GetLevel() TypeUser {
	return el.level
}

func (el *user) SetLevel(value TypeUser) {
	el.level = value
}

func (el *user) IsLogged() bool {
	return true
}
