package server

type ssl struct {
	Enabled                  bool        `json:"enabled"`
	Certificate              string      `json:"certificate"`
	CertificateKey           string      `json:"certificateKey"`
	Version                  sslVersion  `json:"version"`
	CurvePreferences         interface{} `json:"curvePreferences"`
	PreferServerCipherSuites bool        `json:"preferServerCipherSuites"`
	CipherSuites             interface{} `json:"cipherSuites"`
	X509                     sslX509     `json:"x509"`
}
