package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gihub.com/jshiwam/building-microservices-in-go/handlers"
)

func main() {
	// var buf bytes.Buffer
	l := log.New(os.Stdout, "product-api ", log.Ldate|log.Ltime|log.Lshortfile)
	sm := http.NewServeMux()

	l.Println(time.Now(), time.Now().Add(10*time.Second))
	hello := handlers.NewHello(l)
	bye := handlers.NewBye(l)
	sm.Handle("/", hello)
	sm.Handle("/bye", bye)

	_server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		l.Println("calling ListenAndServe")
		serveErr := _server.ListenAndServe()
		l.Println("ListenAndServe called")
		if serveErr != nil {
			l.Fatal("ListenAndServe returns err ", serveErr)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	// signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Terminate signal received, graceful shutdown", sig)
	// _context, _ := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	_context, _ := context.WithTimeout(context.Background(), 30*time.Second)
	shutdownErr := _server.Shutdown(_context)
	if shutdownErr != nil {
		l.Fatal("Server Shutdown err ", shutdownErr)
	}
}
