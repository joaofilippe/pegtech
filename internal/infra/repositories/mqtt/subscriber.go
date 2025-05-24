package mqtt

import (
	"encoding/json"
	"log"
)

// Message representa uma mensagem MQTT
type Message struct {
	Topic   string
	Payload []byte
}

// MessageHandler é uma função que processa mensagens MQTT
type MessageHandler func(topic string, payload []byte)

// Subscriber gerencia as subscrições MQTT
type Subscriber struct {
	client   *MqttClient
	handlers map[string]MessageHandler
	msgChan  chan Message
}

// NewSubscriber cria uma nova instância do Subscriber
func NewSubscriber(client *MqttClient) *Subscriber {
	return &Subscriber{
		client:   client,
		handlers: make(map[string]MessageHandler),
		msgChan:  make(chan Message, 100), // Buffer de 100 mensagens
	}
}

// Subscribe adiciona um handler para um tópico específico
func (s *Subscriber) Subscribe(topic string, handler MessageHandler) error {
	s.handlers[topic] = handler
	return s.client.Subscribe(topic, func(payload []byte) {
		// Envia a mensagem para o canal
		s.msgChan <- Message{
			Topic:   topic,
			Payload: payload,
		}
		// Executa o handler
		handler(topic, payload)
	})
}

// GetMessageChannel retorna o canal de mensagens
func (s *Subscriber) GetMessageChannel() <-chan Message {
	return s.msgChan
}

// SubscribeToPackageRegistration subscreve ao tópico de registro de pacotes
func (s *Subscriber) SubscribeToPackageRegistration() error {
	return s.Subscribe("locker/package/register", func(topic string, payload []byte) {
		// O payload já é uma string, não precisa decodificar JSON
		packageCode := string(payload)
		log.Printf("%s", packageCode)
	})
}

// SubscribeToPackagePickup subscreve ao tópico de retirada de pacotes
func (s *Subscriber) SubscribeToPackagePickup() (chan []byte, error) {
	lockerChan := make(chan []byte, 100) // Buffer de 100 mensagens

	err := s.Subscribe("locker/package/pickup", func(topic string, payload []byte) {
		lockerChan <-payload
	})

	if err != nil {
		close(lockerChan)
		return nil, err
	}

	return lockerChan, nil
}

// SubscribeToLockerAvailable subscreve ao tópico de disponibilidade dos lockers
func (s *Subscriber) SubscribeToLockerAvailable() error {
	type Locker struct {
		LockerID int `json:"locker_id"`
		Ports []int `json:"ports"`
	}

	availableChan := make(chan []Locker, 100) // Buffer de 100 mensagens

	err := s.Subscribe("locker/available", func(topic string, payload []byte) {
		var availableData struct {
			Lockers []Locker `json:"lockers"`
		}

		if err := json.Unmarshal(payload, &availableData); err != nil {
			log.Printf("Erro ao decodificar mensagem MQTT: %v", err)
			return
		}

		// Envia o ID do locker para o canal
		availableChan <- availableData.Lockers
		log.Printf("%v", availableData.Lockers)
	})

	if err != nil {
		close(availableChan)
		return err
	}

	return nil
}

// Start inicia todas as subscrições
func (s *Subscriber) Start() error {
	// Inicia todas as subscrições
	if err := s.SubscribeToPackageRegistration(); err != nil {
		return err
	}

	return nil
}
