package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
)

// initPaths инициализация переменных окружения сертификатов TLS
func initPaths() (string, string, string) {
	caCertPath, ok := os.LookupEnv("CA_CERT_PATH")
	if !ok {
		log.Fatalf("CA_CERT_PATH environment variable not set")
	}
	log.Printf("CA_CERT_PATH=%s", caCertPath)

	certPath, ok := os.LookupEnv("CERT_PATH")
	if !ok {
		log.Fatalf("CERT_PATH environment variable not set")
	}
	log.Printf("CERT_PATH=%s", certPath)

	keyPath, ok := os.LookupEnv("KEY_PATH")
	if !ok {
		log.Fatalf("KEY_PATH environment variable not set")
	}
	log.Printf("KEY_PATH=%s", keyPath)

	return caCertPath, certPath, keyPath
}

// InitTLSConfig функция инициализации TLS конфигурации для MQTT
func InitTLSConfig() (*tls.Config, error) {
	caCertPath, certPath, keyPath := initPaths()

	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("Couldn't load CA certificate: %s", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("Couldn't load client certificate: %s", err)
	}

	return &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: true,
	}, nil
}
