package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/jshiwam/building-microservices-in-go/currency/data"
	protos "github.com/jshiwam/building-microservices-in-go/currency/protos/currency"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	log hclog.Logger
	er  *data.ExchangeRates
}

func NewCurrency(log hclog.Logger, er *data.ExchangeRates) *Currency {
	return &Currency{log: log, er: er}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	rate, err := c.er.GetRate(rr.GetBase().String(), rr.GetDestination().String())

	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
}
