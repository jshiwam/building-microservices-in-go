package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/jshiwam/building-microservices-in-go/product-api/data"
	"github.com/jshiwam/building-microservices-in-go/product-api/handlers"
)

func main() {
	// var buf bytes.Buffer
	l := log.New(os.Stdout, "product-api ", log.Ldate|log.Ltime|log.Lshortfile)
	v := data.NewValidation()
	sm := mux.NewRouter()

	// l.Println(time.Now(), time.Now().Add(10*time.Second))
	hello := handlers.NewHello(l)
	sm.Handle("/", hello)

	bye := handlers.NewBye(l)
	sm.Handle("/bye", bye)

	product := handlers.NewProducts(l, v)
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", product.ListProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", product.GetProductById)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", product.AddProduct)
	postRouter.Use(product.MiddlewareProductValidation)

	updateRouter := sm.Methods(http.MethodPut).Subrouter()
	updateRouter.HandleFunc("/products/{id:[0-9]+}", product.UpdateProduct)
	updateRouter.Use(product.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", product.DeleteProduct)

	opts := middleware.RedocOpts{SpecURL: "./swagger.yaml"}
	docHandler := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", docHandler)

	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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
