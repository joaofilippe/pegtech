package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client mqtt.Client
}

// NewClient cria e conecta um cliente MQTT ao broker externo
func NewClient(broker, clientID, caCertPath string) (*Client, error) {
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler CA: %w", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetTLSConfig(tlsConfig)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Client{client: client}, nil
}

// Publish publica uma mensagem em um tópico
func (c *Client) Publish(topic string, payload []byte) error {
	token := c.client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

// Subscribe subscreve em um tópico
func (c *Client) Subscribe(topic string, handler func([]byte)) error {
	err := c.client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
		handler(msg.Payload())
	})

	if token := c.client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
		handler(msg.Payload())
	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return err.Error()
}

// Disconnect desconecta o cliente
func (c *Client) Disconnect() {
	c.client.Disconnect(250)
}
