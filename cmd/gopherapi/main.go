package main

import (
	"log"
	"net/http"

	"github.com/Zero23ku/helloworld-restapi-go/pkg/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
