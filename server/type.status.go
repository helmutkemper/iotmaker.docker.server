package server

type status struct {
	ExpReg string   `json:"expReg"`
	Value  int      `json:"value"`
	In     []maxMin `json:"in"`
	NotIn  []maxMin `json:"notIn"`
}
