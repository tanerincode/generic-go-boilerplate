package Server

import "github.com/gorilla/mux"

type Api struct {
	addr    string
	handler Handler
}

type Responder struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"result,omitempty"`
}

func New(addr string, svc Handler) *Api {
	return &Api{
		addr:    addr,
		handler: svc,
	}
}

func (a *Api) Start() error {

	router := mux.NewRouter()
	router.HandleFunc("/", a.handler.rootHandler).Methods("GET")

	return nil
}
