package types

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
)

type FileTransferData struct {
	FileName     string    `json:"file_name"`
	SentTo       string    `json:"sent_to"`
	Time         time.Time `json:"time"`
	Size         int64     `json:"size"`
	ReceivedFrom string    `json:"received_from"`
}

type Client struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	TransferData []FileTransferData `json:"transfer_data"`
	Conn         *websocket.Conn    `json:"-"`
	IsOnline     bool               `json:"is_online"`
}

var (
	Connections     = make(map[string]*Client)
	ConnectionMutex sync.RWMutex
)

type PostgresClient struct {
	db *pgx.Conn
}

type PostgresFileTrasferData struct {
	db *pgx.Conn
}
