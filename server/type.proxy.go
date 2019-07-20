package server

import (
	"math/rand"
	"net/http"
)

type proxy struct {
	consecutiveErrors  int
	consecutiveSuccess int
	errors             int
	success            int

	keyProxy  int
	keyServer int

	lastError      error
	lastRoundError bool

	MaxAttemptToRescueLoop int         `yaml:"maxAttemptToRescueLoop"`
	IgnorePort             bool        `yaml:"ignorePort"`
	Host                   string      `yaml:"host"`
	Bind                   []bind      `yaml:"bind"`
	LoadBalancing          string      `yaml:"loadBalancing"`
	Path                   string      `yaml:"path"`
	HealthCheck            healthCheck `yaml:"healthCheck"`
	Servers                []servers   `yaml:"servers"`
}

func (el *proxy) ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {

	//w.WriteHeader(500)
	el.consecutiveErrors += 1
	el.consecutiveSuccess = 0
	el.consecutiveErrors += 1
	el.consecutiveSuccess = 0
	el.errors += 1
	el.lastRoundError = true
	el.lastError = err

	//seelog.Criticalf("1 server host %v error - %v", hostServer, err.Error())
}

func (el *proxy) SuccessHandler(w http.ResponseWriter, r *http.Request, err error) {

	//seelog.Criticalf("1 server host %v error - %v", hostServer, err.Error())
}

func (el *proxy) ModifyResponse(resp *http.Response) error {
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (el *proxy) roundRobin() (string, int) {
	randNumber := rand.Float64()

	for serverKey, serverData := range el.Servers {

		if randNumber <= serverData.Weight {
			return serverData.Host, serverKey
		}

	}

	return "", -1
}

func (el *proxy) random() (string, int) {
	randNumber := rand.Intn(len(el.Servers))
	return el.Servers[randNumber].Host, randNumber
}
