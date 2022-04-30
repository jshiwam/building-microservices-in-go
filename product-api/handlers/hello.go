package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type Hello struct {
	log hclog.Logger
}

func NewHello(l hclog.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(rwriter http.ResponseWriter, request *http.Request) {
	h.log.Info("Hello World")

	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oooops"))

		http.Error(rwriter, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rwriter, "Hello %s", d)
}
