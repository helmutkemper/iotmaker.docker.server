package server

func ExampleNewServer() {

	var project = Project{
		ListenAndServer: ListenAndServer{
			InAddress: "0.0.0.0:8080",
		},
		DebugServerEnable: true,
	}

	// Output:
	//

}
