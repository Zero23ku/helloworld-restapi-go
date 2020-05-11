package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

// Server describe server methods and the return type of such methods
type Server interface {
	Router() http.Handler
}

// New initialice the API and its routes
func New() Server {
	a := &api{}
	r := mux.NewRouter()
	r.HandleFunc("/hello", a.fetchSalute).Methods(http.MethodGet)
	r.HandleFunc("/hello/{NAME:[a-zA-Z]+}", a.fetchSalute).Methods(http.MethodGet)
	a.router = r
	return a
}

// Server interface implementation
func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchSalute(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["NAME"]
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	if len(name) > 0 {
		json.NewEncoder(response).Encode("Hola " + name)

	} else {
		json.NewEncoder(response).Encode("Hola mundo")
	}
	return

}
