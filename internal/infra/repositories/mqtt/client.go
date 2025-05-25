package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	client mqtt.Client
}

// NewClient cria e conecta um cliente MQTT ao broker externo
func NewClient(broker, port, protocol, username, password, clientID, caCertPath string) (*MqttClient, error) {
	connectAddress := fmt.Sprintf("%s://%s:%s", protocol, broker, port)
	tlsConfig := loadTLSConfig(caCertPath)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetClientID(clientID)
	opts.SetKeepAlive(time.Second * 60)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetTLSConfig(tlsConfig)
	opts.SetKeepAlive(time.Second * 60)

	client := mqtt.NewClient(opts)
	token := client.Connect()

	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}

	return &MqttClient{client: client}, nil
}

// Publish publica uma mensagem em um tópico
func (c *MqttClient) Publish(topic string, payload any) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	fmt.Printf("\n\nPublished\nTopic: %s\nPayload: %s\n", topic, string(jsonPayload))

	token := c.client.Publish(topic, 0, false, jsonPayload)
	token.Wait()
	return token.Error()
}

// Subscribe subscreve em um tópico
func (c *MqttClient) Subscribe(topic string, handler func([]byte)) error {
	token := c.client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
		fmt.Printf("\n\nMessage received\nTopic: %s\nData: %s\n", msg.Topic(), string(msg.Payload()))
		handler(msg.Payload())
	})

	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Disconnect desconecta o cliente
func (c *MqttClient) Disconnect() {
	c.client.Disconnect(250)
}

func loadTLSConfig(caFile string) *tls.Config {
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = false
	if caFile != "" {
		certpool := x509.NewCertPool()
		ca, err := os.ReadFile(caFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		certpool.AppendCertsFromPEM(ca)
		tlsConfig.RootCAs = certpool
	}
	return &tlsConfig
}
