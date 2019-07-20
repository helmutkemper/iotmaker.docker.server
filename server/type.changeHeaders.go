package server

type changeHeaders struct {
	Number int      `json:"number"`
	Header []header `json:"header"`
	Rate   float64  `json:"rate"`
}
