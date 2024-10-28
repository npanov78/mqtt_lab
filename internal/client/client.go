package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	server "mqtt_lab.com/internal/server"
	"time"
)

var broker, username, password = server.InitConfig()

// initClientOptions функция инициализации параметров клиента
func initClientOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().AddBroker(broker)

	//tlsConfig, err := server.InitTLSConfig()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//opts.SetTLSConfig(tlsConfig)

	opts.SetClientID("MQTT_Client")

	//opts.SetUsername(username)
	//opts.SetPassword(password)

	opts.OnConnect = onConnect

	return opts
}

// onConnect функция подписки и прослушивания очереди
func onConnect(client mqtt.Client) {
	if token := client.Subscribe(server.Topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Intercepted message: %s\n", msg.Payload())
	}); token.Wait() && token.Error() != nil {
		log.Fatalf("Error: %e", token.Error())
	}
}

// StartClient функция запуска клиента
func StartClient() {
	opts := initClientOptions()

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("Broker %s connection error: %s", broker, token.Error())
		log.Printf("Sleep for 5 seconds...")
		time.Sleep(5 * time.Second)

		return
	}

	log.Printf("Subscribed to topic %s. Listening for messages...", server.Topic)
	select {}
}
