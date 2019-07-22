package server

import (
	"encoding/json"
	"fmt"
	iot_json "github.com/helmutkemper/iotmaker.server.json"
	"github.com/xenolf/lego/log"
	"io/ioutil"
	"net/http"
	"sync"
)

func test(w http.ResponseWriter, r *http.Request) {
	out := iot_json.NewJSonOut()
	out.Objects = []map[string]interface{}{
		{
			"string": "esta vivo",
		},
	}
	w.Write(out.Byte())
}

func ExampleNewServer() {

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
	}

	var wg sync.WaitGroup
	var toOut interface{}

	wg.Add(1)
	go func(project Project) {
		wg.Done()
		log.Fatalf("server error: %v", NewServer(project))
	}(project)

	wg.Wait()

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

	fmt.Printf("%v", toOut.(map[string]interface{})["Objects"])

	// Output:
	// status: 200 OK
	// [map[string:esta vivo]]

}
