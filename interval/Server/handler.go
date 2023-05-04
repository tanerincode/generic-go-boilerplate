package Server

import (
	"encoding/json"
	"log"
	"net/http"
	"tanerincode/generic-go-boilerplate/interval/storage"
)

type Handler interface {
	rootHandler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	storage storage.Storage
}

func NewHandler(store storage.Storage) Handler {
	return &handler{storage: store}
}

func (h *handler) rootHandler(w http.ResponseWriter, r *http.Request) {
	result := &Responder{}

	resp, err := json.Marshal(result)
	if err != nil {
		log.Printf("marshal error %v", err)
	}

	_, _ = w.Write(resp)
}
