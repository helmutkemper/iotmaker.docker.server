package server

import (
	"os"
	"runtime/debug"
	"strings"
)

// pt-br: imprime a pilha de erro caso ocorra um erro do tipo 'multiple response.WriteHeader'
//
// en: prints the error stack if an 'multiple response.WriteHeader' type error occurs
type DebugLogger struct{}

func (d DebugLogger) Write(p []byte) (n int, err error) {
	s := string(p)
	if strings.Contains(s, "multiple response.WriteHeader") {
		debug.PrintStack()
	}
	return os.Stderr.Write(p)
}
