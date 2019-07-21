package server

import "net/http"

type Handle struct {
	Method              string
	Func                func(w http.ResponseWriter, r *http.Request)
	FuncOnError         func(w http.ResponseWriter, r *http.Request)
	FuncOnSecurityError func(w http.ResponseWriter, r *http.Request)
	Security            func(w http.ResponseWriter, r *http.Request) (error, bool)
	HeaderToAdd         map[HeaderList]HeaderApplication
}
