package cert

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"

	"google.golang.org/grpc/credentials"
)

func GetTlsCredentials() (credentials.TransportCredentials, error) {
	clientCert, clientCertErr := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem") // 加载服务器证书
	// handle clientCertErr
	if clientCertErr != nil {
		return nil, clientCertErr
	}
	certPool := x509.NewCertPool() // CA检查的证书池子
	caCert, caCertErr := os.ReadFile("cert/ca-cert.pem")
	if caCertErr != nil {
		return nil, caCertErr
	}

	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("failed to append the CA certs")
	}
	return credentials.NewTLS(
		&tls.Config{
			ClientAuth:   tls.RequireAnyClientCert,
			Certificates: []tls.Certificate{clientCert},
			ClientCAs:    certPool,
		}), nil
}
