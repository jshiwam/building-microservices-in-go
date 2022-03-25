package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	_log *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(rwriter http.ResponseWriter, request *http.Request) {
	h._log.Println("Hello World")

	d, err := ioutil.ReadAll(request.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oooops"))

		http.Error(rwriter, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rwriter, "Hello %s", d)
}
