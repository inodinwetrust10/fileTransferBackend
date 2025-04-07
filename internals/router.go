package internals

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	Router   = http.NewServeMux()
	upgrader = websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
	}
)

func InitRoutes() {
	Router.HandleFunc("/ws", HandleConn)
}
