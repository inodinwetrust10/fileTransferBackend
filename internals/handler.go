package internals

import (
	"net/http"

	"github.com/inodinwetrust10/filetransfer/utils"
)

func HandleConn(w http.ResponseWriter, r *http.Request) {
	Conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, map[string]string{
			"error": "websocket connection failed",
			"cause": err.Error(),
		})
		return
	}
	defer Conn.Close()
}
