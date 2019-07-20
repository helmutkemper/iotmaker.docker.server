package server

type ssl struct {
	Enabled                  bool        `yaml:"enabled"`
	Certificate              string      `yaml:"certificate"`
	CertificateKey           string      `yaml:"certificateKey"`
	Version                  sslVersion  `yaml:"version"`
	CurvePreferences         interface{} `yaml:"curvePreferences"`
	PreferServerCipherSuites bool        `yaml:"preferServerCipherSuites"`
	CipherSuites             interface{} `yaml:"cipherSuites"`
	X509                     sslX509     `yaml:"x509"`
}
