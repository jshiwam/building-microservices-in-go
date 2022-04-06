package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	protos "github.com/jshiwam/building-microservices-in-go/currency/protos/currency"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(log hclog.Logger) *Currency {
	return &Currency{log: log}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
