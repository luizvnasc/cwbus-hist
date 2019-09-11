package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const port = "8081"

func TestServer(t *testing.T) {

	t.Run("Iniciando servidor sem informar a porta", func(t *testing.T) {
		server := New("")
		defer server.Shutdown()
		quit := make(chan bool)
		defer close(quit)
		go func() {
			go server.Run()
			for {
				select {
				case <-quit:
					log.Println("Finalizando servidor")
					return
				}
			}
		}()
		// Espera o servidor subir
		time.Sleep(30 * time.Millisecond)
		_, err := net.Dial("tcp", ":"+port)
		if err != nil {
			t.Errorf("Site unreachable, error: %q", err)
		}
		quit <- true
	})

	t.Run(fmt.Sprintf("Inciando servidor na porta %q", port), func(t *testing.T) {
		server := New(port)
		defer server.Shutdown()
		quit := make(chan bool)
		defer close(quit)
		go func() {
			go server.Run()
			for {
				select {
				case <-quit:
					log.Println("Finalizando servidor")
					return
				}
			}
		}()
		// Espera o servidor subir
		time.Sleep(30 * time.Millisecond)
		_, err := net.Dial("tcp", ":"+port)
		if err != nil {
			t.Errorf("Site unreachable, error: %q", err)
		}
		quit <- true

	})

	testCases := []struct {
		path   string
		method string
		body   io.Reader
		want   string
	}{
		{"/versao", http.MethodGet, nil, "0.0.1"},
	}
	server := New("")
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testando rota %q", tc.path), func(t *testing.T) {
			req, _ := http.NewRequest(tc.method, tc.path, tc.body)
			res := httptest.NewRecorder()

			http.HandlerFunc(server.routes[tc.path].Handler()).ServeHTTP(res, req)

			got := res.Body.String()
			if got != tc.want {
				t.Errorf("Erro ao testar rota %q. Resposta esperada: %q, obtida: %q", tc.path, tc.want, got)
			}
		})
	}

}
