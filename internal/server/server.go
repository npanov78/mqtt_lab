package server

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Printf("Connected to %s", broker)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Disconnected from %s", broker)
}

var broker, username, password = InitConfig()

// generateSensorData функция генерирует данные для сенсора
func generateSensorData() string {
	temperature := 20.0 + rand.Float64()*5.0 // от 20.0 до 25.0
	humidity := 30.0 + rand.Float64()*30.0   // от 30.0 до 60.0
	return fmt.Sprintf("temperature: %.2f, humidity: %.2f", temperature, humidity)
}

// initBrokerOptions функция генерирует параметры подключения к брокеру
func initBrokerOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().AddBroker(broker)

	//tlsConfig, err := InitTLSConfig()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//opts.SetTLSConfig(tlsConfig)

	opts.SetClientID("MQTT_Server")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	//opts.SetUsername(username)
	//opts.SetPassword(password)

	return opts
}

// StartServer функция запуска сервера
func StartServer() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	opts := initBrokerOptions()

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("Broker %s connection error: %s", broker, token.Error())
		log.Printf("Sleep for 5 seconds...")
		time.Sleep(5 * time.Second)
		return
	}

	log.Printf("MQTT server started, publishing data...")

	for {
		payload := generateSensorData()
		token := client.Publish(Topic, 0, false, payload)
		token.Wait()
		log.Printf("Published: %s\n", payload)
		time.Sleep(5 * time.Second)
	}
}
