package server

type HeaderApplication int

func (el HeaderApplication) String() string {
	return HeaderApplications[el]
}

var HeaderApplications = [...]string{
	"",
	"application/json",
}

const (
	KHeaderApplicationTypeJSon HeaderApplication = iota + 1
)
