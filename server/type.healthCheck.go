package server

type healthCheck struct {
	Enabled         bool  `json:"enabled"`
	Interval        int   `json:"interval"`
	Fails           int   `json:"fails"`
	Passes          int   `json:"passes"`
	Uri             int   `json:"rui"`
	SuspendInterval int   `json:"suspendInterval"`
	Match           match `json:"match"`
}
