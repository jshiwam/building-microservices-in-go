package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/jshiwam/building-microservices-in-go/currency/data"
	protos "github.com/jshiwam/building-microservices-in-go/currency/protos/currency"
	"github.com/jshiwam/building-microservices-in-go/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	log.Info("Starting server")
	gs := grpc.NewServer()
	er, err := data.NewRates(log)

	if err != nil {
		log.Error("Unable to get rates", "error", err)
		os.Exit(1)
	}
	c := server.NewCurrency(log, er)

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
