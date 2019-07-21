package server

import (
	iot_json "github.com/helmutkemper/iotmaker.server.json"
	"github.com/xenolf/lego/log"
	"net/http"
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

	log.Fatalf("server error: %v", NewServer(project))

	// Output:
	//

}
