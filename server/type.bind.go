package server

type bind struct {
	Host       string `yaml:"host"`
	IgnorePort bool   `yaml:"ignorePort"`
}
