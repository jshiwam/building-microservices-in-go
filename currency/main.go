package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	protos "github.com/jshiwam/building-microservices-in-go/currency/protos/currency"
	"github.com/jshiwam/building-microservices-in-go/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	log := hclog.Default()
	log.Info("Starting server")
	gs := grpc.NewServer()

	c := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, c)

	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))

	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	log.Info("Listening on port 9092")
	gs.Serve(l)
}
