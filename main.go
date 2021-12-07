package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"xiaolong.com/v2/handlers"
)

func main() {
	l := log.New(os.Stdout, "service-api", log.LstdFlags)
	l.Println("logger initated")
	aph := handlers.NewArtPiece(l)
	bh := handlers.NewBye(l)
	sm := http.NewServeMux()

	sm.Handle("/", aph)
	sm.Handle("/bye", bh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	l.Println("server starting...")
	go func() {
		l.Println("serving HTTP...")
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("gracefully shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
