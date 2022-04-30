package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	l     hclog.Logger
	rates map[string]float64
}

func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{l: l, rates: map[string]float64{}}
	err := er.getRates()

	if err != nil {
		return nil, err
	}
	return er, nil
}

func (er *ExchangeRates) GetRate(base, dest string) (float64, error) {
	br, ok := er.rates[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", base)
	}

	dr, ok := er.rates[dest]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", dest)
	}
	return dr / br, nil
}
func (er *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		er.l.Error("[Error] Cannot Get Exchange Rates", "error", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected error code 200 got %d", resp.StatusCode)
	}
	md := &Cubes{}
	err = xml.NewDecoder(resp.Body).Decode(md)

	if err != nil {
		er.l.Error("[Error] Cannot Decode XML", "error", err)
		return err
	}
	for _, cube := range md.CubeData {
		rate, err := strconv.ParseFloat(cube.Rate, 64)
		if err != nil {
			er.l.Error("[Error] Cannot convert rate from string to float", "error", err)
		}
		er.rates[cube.Currency] = rate
	}
	er.rates["EUR"] = 1
	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
