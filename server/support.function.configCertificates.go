package server

import (
	"crypto/tls"
	"errors"
	"net/http"
	"os"
)

// pt-br: configura os certificados ssl
//
// en: configures ssl certificates
func configCertificates(config ssl, server *http.Server) error {
	var err error

	if config.Enabled == true {

		var certificatesList = tls.Certificate{}

		if config.X509.Certificate != "" && config.X509.CertificateKey != "" {

			if _, err = os.Stat(config.X509.Certificate); os.IsNotExist(err) {
				return errors.New("sll x509 certificate error: " + err.Error())
			}

			if _, err = os.Stat(config.X509.CertificateKey); os.IsNotExist(err) {
				return errors.New("sll x509 certificate key error: " + err.Error())
			}

			certificatesList, err = tls.LoadX509KeyPair(config.X509.Certificate, config.X509.CertificateKey)
			if err != nil {
				return errors.New("sll x509 certificate load pair error: " + err.Error())
			}

		}

		var tlsMinVersion uint16 = 0
		if config.Version.Min == 10 {
			tlsMinVersion = tls.VersionTLS10
		} else if config.Version.Min == 11 {
			tlsMinVersion = tls.VersionTLS11
		} else if config.Version.Min == 12 {
			tlsMinVersion = tls.VersionTLS12
		} else if config.Version.Min == 30 {
			tlsMinVersion = tls.VersionSSL30
		}

		var tlsMaxVersion uint16 = 0
		if config.Version.Max == 10 {
			tlsMaxVersion = tls.VersionTLS10
		} else if config.Version.Max == 11 {
			tlsMaxVersion = tls.VersionTLS11
		} else if config.Version.Max == 12 {
			tlsMaxVersion = tls.VersionTLS12
		} else if config.Version.Max == 30 {
			tlsMaxVersion = tls.VersionSSL30
		}

		var curveIdList = make([]tls.CurveID, len(config.CurvePreferences.([]string)))
		for k, v := range config.CurvePreferences.([]string) {
			if v == "P256" {
				curveIdList[k] = tls.CurveP256
			} else if v == "P384" {
				curveIdList[k] = tls.CurveP384
			} else if v == "P521" {
				curveIdList[k] = tls.CurveP521
			} else if v == "X25519" {
				curveIdList[k] = tls.X25519
			}
		}

		var cipherSuitesList = make([]uint16, len(config.CipherSuites.([]string)))
		for k, v := range config.CurvePreferences.([]string) {
			if v == "TLS_RSA_WITH_RC4_128_SHA" {
				cipherSuitesList[k] = tls.TLS_RSA_WITH_RC4_128_SHA
			} else if v == "TLS_RSA_WITH_3DES_EDE_CBC_SHA" {
				cipherSuitesList[k] = tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA
			} else if v == "TLS_RSA_WITH_AES_128_CBC_SHA" {
				cipherSuitesList[k] = tls.TLS_RSA_WITH_AES_128_CBC_SHA
			} else if v == "TLS_RSA_WITH_AES_256_CBC_SHA" {
				cipherSuitesList[k] = tls.TLS_RSA_WITH_AES_256_CBC_SHA
			} else if v == "TLS_RSA_WITH_AES_128_CBC_SHA256" {
				cipherSuitesList[k] = tls.TLS_RSA_WITH_AES_128_CBC_SHA256
			} else if v == "TLS_RSA_WITH_AES_128_GCM_SHA256" {
				cipherSuitesList[k] = tls.TLS_RSA_WITH_AES_128_GCM_SHA256
			} else if v == "TLS_RSA_WITH_AES_256_GCM_SHA384" {
				cipherSuitesList[k] = tls.TLS_RSA_WITH_AES_256_GCM_SHA384
			} else if v == "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA" {
				cipherSuitesList[k] = tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA
			} else if v == "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA" {
				cipherSuitesList[k] = tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA
			} else if v == "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA" {
				cipherSuitesList[k] = tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA
			} else if v == "TLS_ECDHE_RSA_WITH_RC4_128_SHA" {
				cipherSuitesList[k] = tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA
			} else if v == "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA" {
				cipherSuitesList[k] = tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA
			} else if v == "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA" {
				cipherSuitesList[k] = tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA
			} else if v == "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA" {
				cipherSuitesList[k] = tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA
			} else if v == "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256" {
				cipherSuitesList[k] = tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256
			} else if v == "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256" {
				cipherSuitesList[k] = tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256
			} else if v == "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256" {
				cipherSuitesList[k] = tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
			} else if v == "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256" {
				cipherSuitesList[k] = tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
			} else if v == "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384" {
				cipherSuitesList[k] = tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
			} else if v == "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384" {
				cipherSuitesList[k] = tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
			} else if v == "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305" {
				cipherSuitesList[k] = tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305
			} else if v == "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305" {
				cipherSuitesList[k] = tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305
			} else if v == "TLS_FALLBACK_SCSV" {
				cipherSuitesList[k] = tls.TLS_FALLBACK_SCSV
			}
		}

		server.TLSConfig = &tls.Config{
			MinVersion:               tlsMinVersion,
			MaxVersion:               tlsMaxVersion,
			CurvePreferences:         curveIdList,
			PreferServerCipherSuites: config.PreferServerCipherSuites,
			CipherSuites:             cipherSuitesList,
			Certificates:             []tls.Certificate{certificatesList},
		}

	}

	return nil
}
