package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inodinwetrust10/filetransfer/internals"
	"github.com/inodinwetrust10/filetransfer/internals/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := database.ConnectDB()
	defer db.Close(context.Background())
	internals.InitRoutes(db)

	server := &http.Server{
		Handler: internals.Router,
		Addr:    "localhost:8080",
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	go func() {
		log.Println("Server started")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("error starting the server:", err)
		}
		if err == http.ErrServerClosed {
			log.Println("Server stopped accepting new connections")
		}
	}()
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown the server")
	} else {
		slog.Info("server shutdown successfully")
	}
}
