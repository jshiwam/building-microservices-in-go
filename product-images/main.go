package main

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/jshiwam/building-microservices-in-go/product-images/files"
	"github.com/jshiwam/building-microservices-in-go/product-images/handlers"
)

var basePath = "./imagestore"
var bindAddress = ":9091"
var logLevel = "debug"

func main() {
	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString(logLevel),
		},
	)

	// create a logger for the server from the default logger
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	stor, err := files.NewLocal(basePath, 1024*1000*5)

	if err != nil {
		l.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	// create the handlers
	fh := handlers.NewFiles(stor, l)
	mw := handlers.GzipHandler{}
	// create a new servemux and register the handlers
	sm := mux.NewRouter()

	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

	// get files

	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))))
	gh.Use(mw.GzipMiddleware)
	// create a new server
	s := http.Server{
		Addr:         bindAddress,
		Handler:      sm,
		ErrorLog:     sl,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Info("Starting server", "bind_address", bindAddress)
		err := s.ListenAndServe()
		if err != nil {
			l.Error("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received
	<-c

}
