package server

type servers struct {
	consecutiveErrors  int
	consecutiveSuccess int
	errors             int
	success            int

	lastRoundError bool

	Host     string  `json:"host"`
	Weight   float64 `json:"weight"`
	OverLoad int     `json:"overLoad"`
}
