package pegtech

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	mqttserver "github.com/mochi-mqtt/server/v2"
	auth "github.com/mochi-mqtt/server/v2/hooks/auth"
	listeners "github.com/mochi-mqtt/server/v2/listeners"
)

func startMQTT() {
	// Cria uma instância do broker MQTT
	mqttServer := mqttserver.New(nil)

	// Permite todas as conexões (apenas para desenvolvimento)
	_ = mqttServer.AddHook(new(auth.AllowHook), nil)

	// Cria um listener TCP na porta 1883
	tcp := listeners.NewTCP(listeners.Config{ID: "t1", Address: ":1883"})
	err := mqttServer.AddListener(tcp)
	if err != nil {
		log.Fatalf("Erro ao adicionar listener TCP: %v", err)
	}

	go func() {
		log.Println("Iniciando servidor MQTT na porta :1883...")
		err := mqttServer.Serve()
		if err != nil {
			log.Fatalf("Erro ao iniciar o servidor MQTT: %v", err)
		}
	}()
}

func main() {
	// Inicia o servidor MQTT em uma goroutine
	startMQTT()
	// Create a new Echo instance
	e := echo.New()

	// Define a route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to PegTech API!")
	})

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
