package server

type pygocentrus struct {
	Enabled          bool            `json:"enabled"`
	Delay            rateMaxMin      `json:"delay"`
	DontRespond      rateMaxMin      `json:"dontRespond"`
	ChangeLength     float64         `json:"changeLength"`
	ChangeContent    changeContent   `json:"changeContent"`
	DeleteContent    float64         `json:"deleteContent"`
	ChangeHeaders    []changeHeaders `json:"changeHeaders"`
	successfulAttack bool            `json:"-"`
}

func (el *pygocentrus) SetAttack() {
	el.successfulAttack = true
}

func (el *pygocentrus) GetAttack() bool {
	return el.successfulAttack
}

func (el *pygocentrus) ClearAttack() {
	el.successfulAttack = false
}
