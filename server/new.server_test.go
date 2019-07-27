package server

import (
	"encoding/json"
	"fmt"
	iotJson "github.com/helmutkemper/iotmaker.server.json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func test(w http.ResponseWriter, _ *http.Request) {
	out := iotJson.NewJSonOut()
	out.Objects = []map[string]interface{}{
		{
			"string": "esta vivo",
		},
	}

	_, err := w.Write(out.Byte())
	if err != nil {
		log.Fatalf("http server error: %v", err.Error())
	}
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	out := iotJson.NewJSonOut()
	out.Meta.AddError("page not found")

	_, err := w.Write(out.Byte())
	if err != nil {
		log.Fatalf("http server error: %v", err.Error())
	}
}

func testPageFound() {

	var toOut interface{}

	r, err := http.Get("http://0.0.0.0:8080/test")
	if err != nil {
		log.Fatalf("test error: %v", err.Error())
		return
	}

	fmt.Printf("status: %v\n", r.Status)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("test error: %v", err.Error())
		return
	}

	err = json.Unmarshal(body, &toOut)
	if err != nil {
		log.Fatalf("test error: %v", err.Error())
		return
	}

	fmt.Printf("Test page found:\n%v\n", toOut.(map[string]interface{})["Objects"])
}

func testPageNotFound() {

	var toOut interface{}

	r, err := http.Get("http://0.0.0.0:8080/notFoundPage")
	if err != nil {
		log.Fatalf("test error: %v", err.Error())
		return
	}

	fmt.Printf("status: %v\n", r.Status)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("test error: %v", err.Error())
		return
	}

	err = json.Unmarshal(body, &toOut)
	if err != nil {
		log.Fatalf("test error: %v", err.Error())
		return
	}

	fmt.Printf("Test page not found:\n%v\n", toOut.(map[string]interface{})["Meta"].(map[string]interface{})["Error"].([]interface{})[0])
}

func ExampleNewServer() {

	// pt-br: configuração básica do novo servidor
	// en: basic server configuration
	var project = Project{
		ListenAndServer: ListenAndServer{
			InAddress: "0.0.0.0:8080",
		},
		DebugServerEnable: true,
		Handle: map[string]Handle{
			"/test": {
				Func: test,
				HeaderToAdd: map[HeaderList]HeaderApplication{
					KHeaderListContentType: KHeaderApplicationTypeJSon,
				},
				Method: http.MethodGet,
			},
		},
		FuncPageNotFound: notFound,
	}

	// pt-br: espera o servidor subir
	// en: wait for goes up
	var wg sync.WaitGroup
	wg.Add(1)

	// pt-br: go runtime permite que vários servidores possam existir no mesmo código apenas trocando a porta de acesso
	// en: go runtime allows multiple servers to exist in the same code just by changing the access port
	go func(project Project) {
		wg.Done()
		log.Fatalf("server error: %v", NewServer(project))
	}(project)

	wg.Wait()

	testPageFound()
	testPageNotFound()

	// Output:
	// status: 200 OK
	// Test page found:
	// [map[string:esta vivo]]
	// status: 200 OK
	// Test page not found:
	// page not found

}
