package server

type HeaderList int

func (el HeaderList) String() string {
	return HeaderLists[el]
}

var HeaderLists = [...]string{
	"",
	"Content-Type",
}

const (
	KHeaderListContentType HeaderList = iota + 1
)
