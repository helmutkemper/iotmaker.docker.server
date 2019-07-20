package server

import (
	"errors"
	"log"
	"net/http"
	"os"
)

// pt-br: Monta um servidor http e deve ser montada dentro de um go runtime
//
// Exemplo: go func(config Project){ NewServer( ... ) }(config...)
//
// en: Mounts an http server and must be mounted within a go runtime
//
// Example: go func(config Project){ NewServer( ... ) }(config...)
func NewServer(config Project) error {
	var err error

	server := http.NewServeMux()

	for _, staticPath := range config.Static {

		if _, err = os.Stat(staticPath.FilePath); os.IsNotExist(err) {
			err = errors.New("static dir error: " + err.Error())
			return err
		}

		server.Handle("/"+staticPath.ServerPath+"/", http.StripPrefix("/"+staticPath.ServerPath+"/", http.FileServer(http.Dir(staticPath.FilePath))))
	}

	server.HandleFunc("/", config.HandleFunc)

	newServer := &http.Server{
		//TLSNextProto:               make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		Addr:    config.ListenAndServer.InAddress,
		Handler: server,
	}

	if config.DebugServerEnable == true {
		newServer.ErrorLog = log.New(DebugLogger{}, "", 0)
	}

	if err = configCertificates(config.Sll, newServer); err != nil {
		return err
	}

	if config.Sll.Certificate != "" && config.Sll.CertificateKey != "" {
		return newServer.ListenAndServeTLS(config.Sll.Certificate, config.Sll.CertificateKey)
	}

	return newServer.ListenAndServe()
}
