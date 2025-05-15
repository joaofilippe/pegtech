package config

import (
	"fmt"
	"os"
)

// MQTTConfig contém as configurações necessárias para conexão MQTT
type MQTTConfig struct {
	Broker     string
	ClientID   string
	CACertPath string
}

// LoadMQTTConfig carrega as configurações MQTT das variáveis de ambiente
func LoadMQTTConfig() (*MQTTConfig, error) {
	broker := os.Getenv("MQTT_BROKER")
	if broker == "" {
		return nil, fmt.Errorf("MQTT_BROKER não configurado")
	}

	clientID := os.Getenv("MQTT_CLIENT_ID")
	if clientID == "" {
		return nil, fmt.Errorf("MQTT_CLIENT_ID não configurado")
	}

	caCertPath := os.Getenv("MQTT_CA_CERT_PATH")
	if caCertPath == "" {
		return nil, fmt.Errorf("MQTT_CA_CERT_PATH não configurado")
	}

	return &MQTTConfig{
		Broker:     broker,
		ClientID:   clientID,
		CACertPath: caCertPath,
	}, nil
}
