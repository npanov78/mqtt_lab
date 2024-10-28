package test

import (
	"fmt"
	"math/rand"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	broker   = "tcp://localhost:1883" // URL вашего MQTT брокера
	topic    = "iot/sensor_data"
	username = "testuser"
	password = "password123"
)

// Функция для генерации случайных данных датчика
func generateSensorData() string {
	temperature := 20.0 + rand.Float64()*5.0 // от 20.0 до 25.0
	humidity := 30.0 + rand.Float64()*30.0   // от 30.0 до 60.0
	return fmt.Sprintf("temperature: %.2f, humidity: %.2f", temperature, humidity)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetClientID("SecureMQTT_Server")

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Secure MQTT server started with credentials, publishing data...")

	for {
		payload := generateSensorData()
		token := client.Publish(topic, 0, false, payload)
		token.Wait()
		fmt.Printf("Published: %s\n", payload)
		time.Sleep(5 * time.Second) // Пауза 5 секунд
	}
}
