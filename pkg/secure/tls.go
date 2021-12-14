package secure

import (
	"crypto/tls"
	"crypto/x509"
)

func GetTlsConfig(c *Certs, certPool *x509.CertPool, protos []string) *tls.Config {
	return &tls.Config{
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{c.CertBytes},
				PrivateKey:  c.Key,
			},
		},
		RootCAs:    certPool,
		ServerName: "localhost",
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
		NextProtos: protos,
	}
}
