package user

type TypeUser int

var TypeUsers = [...]string{
	"",
	"god",
}

func (el TypeUser) String() string {
	return TypeUsers[el]
}

const (
	KUserTypeGod TypeUser = iota + 1
)
