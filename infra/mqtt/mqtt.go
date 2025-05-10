package mqtt

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
)

type MQTTServer struct {
	server *mqtt.Server
}

func NewMQTTServer() *MQTTServer {
	// Create the new MQTT Server.
	server := mqtt.New(&mqtt.Options{
		InlineClient: true, // Enable the inline client for direct publishing
	})

	// Allow all connections.
	server.AddHook(new(auth.AllowHook), nil)

	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP(listeners.Config{
		ID:      "t1",
		Address: ":1883",
	})
	err := server.AddListener(tcp)
	if err != nil {
		panic(err)
	}

	return &MQTTServer{
		server: server,
	}
}

func (s *MQTTServer) Start() error {
	return s.server.Serve()
}

func (s *MQTTServer) Shutdown() error {
	s.server.Close()
	return nil
}

func (s *MQTTServer) Publish(topic string, payload []byte) error {
	return s.server.Publish(topic, payload, false, 0)
}

func (s *MQTTServer) Subscribe(topic string, handler func(topic string, payload []byte)) error {
	return s.server.Subscribe(topic, 0, func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		handler(pk.TopicName, pk.Payload)
	})
}
