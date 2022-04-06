package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bye struct {
	_log *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(rwriter http.ResponseWriter, request *http.Request) {
	b._log.Println("Bye World")
	d, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(rwriter, "Bye Crashed", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rwriter, "Bye %s", d)
}
