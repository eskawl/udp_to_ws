package udp_to_ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Handlers interface {
	UdpWebSocket(c *websocket.Conn, w http.ResponseWriter, r *http.Request)
}

func StartServer(handlers Handlers, options ...Option) {
	serverConfig := &ServerConfig{}

	options = append([]Option{
		WithDefaults(),
	}, options...)

	for _, option := range options {
		option(serverConfig)
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/udp", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println("upgrade:", err)
			return
		}

		defer c.Close()

		handlers.UdpWebSocket(c, w, r)
	})

	log.Printf("Starting WS server on %v", serverConfig.addr)
	http.ListenAndServe(serverConfig.addr, nil)
}

type ServerConfig struct {
	addr string
}

type Option func(*ServerConfig)

func WithAddr(addr string) Option {
	return func(so *ServerConfig) {
		so.addr = addr
	}
}

func WithDefaults() Option {
	return func(so *ServerConfig) {
		if so.addr == "" {
			so.addr = "localhost:8080"
		}
	}
}
