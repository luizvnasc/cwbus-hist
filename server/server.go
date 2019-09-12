// Package server contém a implementação do servidor web da aplicação.
package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/luizvnasc/cwbus-hist/server/router"
)

// CwbusServer é o servidor da aplicação web
type CwbusServer struct {
	addr   string
	routes []*router.Route
	server *http.Server
}

// Run inicia o servidor
func (s *CwbusServer) Run() {

	s.server.ListenAndServe()
}

// Shutdown desliga o servidor
func (s *CwbusServer) Shutdown() {
	// Limpa os handlers registrados antes de desligar o servidor
	http.DefaultServeMux = http.NewServeMux()

	if err := s.server.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}

// New constrói um servidor
func New(port string) *CwbusServer {
	if len(port) == 0 {
		port = "8081"
	}
	server := &http.Server{Addr: ":" + port}
	for _, route := range router.Routes() {
		http.HandleFunc(route.Path(), route.Handler())
	}

	return &CwbusServer{addr: fmt.Sprintf(":%s", port), server: server, routes: router.Routes()}
}
