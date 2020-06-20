package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
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

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		// Error closing or starting listener
		log.Fatalf("HTTP Server listen: %v", err)
	}

	log.Printf("Server is listening on %s", srv.Addr)
	<-idleConnectionsClosed
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Print("handler started")
	defer log.Print("handler ended")

	select {
	case <-time.After(time.Second * 3):
		fmt.Fprintf(w, "hola techcrunch!\n")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// time.Sleep(time.Second * 3)
}
