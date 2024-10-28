package server

import (
	"fmt"
	"os"
)

const (
	Topic = "iot/sensor_data"
)

func InitConfig() (string, string, string) {
	brokerHost, ok := os.LookupEnv("BROKER_HOST")
	if !ok {
		brokerHost = "localhost"
	}

	brokerPort, ok := os.LookupEnv("BROKER_PORT")
	if !ok {
		brokerPort = "1883"
	}

	username, ok := os.LookupEnv("BROKER_USER")
	if !ok {
		username = "user"
	}

	password, ok := os.LookupEnv("BROKER_PASS")
	if !ok {
		password = "password"
	}

	if brokerPort == "1883" {
		return fmt.Sprintf("tcp://%s:%s", brokerHost, brokerPort), username, password
	} else {
		return fmt.Sprintf("ssl://%s:%s", brokerHost, brokerPort), username, password
	}
}
