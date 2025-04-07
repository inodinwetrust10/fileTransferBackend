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
)

func main() {
	internals.InitRoutes()
	server := &http.Server{
		Handler: internals.Router,
		Addr:    "localhost:8080",
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("error starting the server")
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown the server")
	} else {
		slog.Info("server shutdown successfully")
	}
}
