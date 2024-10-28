package test

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	broker   = "tcp://localhost:1883" // URL вашего MQTT брокера
	topic    = "iot/sensor_data"
	username = "testuser"
	password = "password123"
)

func main() {
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetClientID("SecureMQTT_Client")

	// Назначение функции обратного вызова для обработки сообщений
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, func(client MQTT.Client, msg MQTT.Message) {
			fmt.Printf("Intercepted message: %s\n", msg.Payload())
		}); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscribed to topic with credentials. Listening for messages...")
	select {} // Бесконечный цикл для удержания клиента активным
}
