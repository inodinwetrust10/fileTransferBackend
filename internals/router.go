package internals

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
)

var (
	Router   = http.NewServeMux()
	upgrader = websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
	}
)

func InitRoutes(conn *pgx.Conn) {
	Router.HandleFunc("/ws", HandleConn)
}
