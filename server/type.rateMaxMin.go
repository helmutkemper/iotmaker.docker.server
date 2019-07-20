package server

type rateMaxMin struct {
	Rate float64 `json:"rate"`
	Min  int     `json:"min"`
	Max  int     `json:"max"`
}
