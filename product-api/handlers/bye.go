package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type Bye struct {
	log hclog.Logger
}

func NewBye(l hclog.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(rwriter http.ResponseWriter, request *http.Request) {
	b.log.Info("Bye World")
	d, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(rwriter, "Bye Crashed", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rwriter, "Bye %s", d)
}
