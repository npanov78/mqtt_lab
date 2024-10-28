package server

import (
	"fmt"
	"os"
)

const (
	Topic = "iot/sensor_data"
)

func InitConfig() string {
	brokerHost, ok := os.LookupEnv("BROKER_HOST")
	if !ok {
		brokerHost = "localhost"
	}

	brokerPort, ok := os.LookupEnv("BROKER_PORT")
	if !ok {
		brokerPort = "1883"
	}
	return fmt.Sprintf("tcp://%s:%s", brokerHost, brokerPort)
}
