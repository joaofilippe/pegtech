package mqtt

import (
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
		log.Printf("Pacote registrado: %s", packageCode)
	})
}

// Start inicia todas as subscrições
func (s *Subscriber) Start() error {
	// Adicione aqui outras subscrições conforme necessário
	return s.SubscribeToPackageRegistration()
}
