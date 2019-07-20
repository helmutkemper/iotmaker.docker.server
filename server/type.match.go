package server

type match struct {
	Status []status `json:"status"`
	Header []header `json:"header"`
	Body   []string `json:"body"`
}
