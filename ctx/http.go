package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func mains() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	srv := &http.Server{
		Addr:    ":4422",
		Handler: mux,
	}

	idleConnectionsClosed := make(chan struct{})
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch

		// Received interrupt signal. Shutting down...
		if err := srv.Shutdown(context.Background()); err != nil {
			// Erro from closing listeners or context timeout
			log.Printf("Server shutdown: %v\n", err)
		}
		close(idleConnectionsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error closing or starting listener
		log.Fatalf("HTTP Server listen: %v", err)
	}

	<-idleConnectionsClosed
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	r.Context()
}
