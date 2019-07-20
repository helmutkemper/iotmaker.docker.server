package server

import "net/http"

type pygocentrusFunc func(req *http.Request) (resp *http.Response, err error)
