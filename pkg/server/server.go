package serve

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

const (
	timeout = time.Second * 20
)

// Run creates a new routed Runr and runs it.
func Run(port uint16) {
	r := mux.NewRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		Handler:      r,
		IdleTimeout:  timeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("info: gracefully shutting down server...")
	os.Exit(0)
}
