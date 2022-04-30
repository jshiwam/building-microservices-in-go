package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/jshiwam/building-microservices-in-go/product-images/files"
)

type Files struct {
	log   hclog.Logger
	store files.Storage
}

func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, rw, r)
}

func (f *Files) saveFile(id string, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("Save File", id, path)
	fp := filepath.Join(id, path)
	f.store.Save(fp, r.Body)
}
