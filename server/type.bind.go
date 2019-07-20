package server

type bind struct {
	Host       string `json:"host"`
	IgnorePort bool   `json:"ignorePort"`
}
