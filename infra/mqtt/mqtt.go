package mqtt

import (
	"encoding/json"
	"fmt"

	"github.com/joaofilippe/pegtech/application/services"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
)

type MQTTServer struct {
	server        *mqtt.Server
	lockerService *services.LockerService
}

type LockerCommand struct {
	Action   string `json:"action"` // "open" or "status"
	LockerID string `json:"locker_id"`
	Password string `json:"password,omitempty"`
}

func NewMQTTServer(lockerService *services.LockerService) *MQTTServer {
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

	mqttServer := &MQTTServer{
		server:        server,
		lockerService: lockerService,
	}

	// Subscribe to locker commands
	mqttServer.Subscribe("locker/+/command", mqttServer.handleLockerCommand)

	return mqttServer
}

func (s *MQTTServer) handleLockerCommand(topic string, payload []byte) {
	var cmd LockerCommand
	if err := json.Unmarshal(payload, &cmd); err != nil {
		fmt.Printf("Error parsing locker command: %v\n", err)
		return
	}

	switch cmd.Action {
	case "open":
		if err := s.lockerService.OpenLocker(cmd.LockerID, cmd.Password); err != nil {
			fmt.Printf("Error opening locker: %v\n", err)
			return
		}
		// Publish confirmation
		s.Publish(fmt.Sprintf("locker/%s/status", cmd.LockerID), []byte(`{"status":"opened"}`))
	case "status":
		// Handle status update from locker
		// This would need to be implemented in the LockerService
		fmt.Printf("Received status update for locker %s\n", cmd.LockerID)
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
