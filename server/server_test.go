package server

import (
	"fmt"
	"io"
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
		go func() {
			go server.Run()
			for {
				select {
				case <-quit:
					return
				}
			}
		}()
		timeout := time.Duration(1 * time.Second)
		_, err := net.DialTimeout("tcp", "localhost:"+port, timeout)
		if err != nil {
			t.Errorf("Site unreachable, error: %q", err)
		}
		quit <- true
		close(quit)
	})

	server := New(port)
	t.Run(fmt.Sprintf("Inciando servidor na porta %q", port), func(t *testing.T) {
		go server.Run()
		defer server.Shutdown()
		timeout := time.Duration(1 * time.Second)
		_, err := net.DialTimeout("tcp", "localhost:"+port, timeout)
		if err != nil {
			t.Errorf("Site unreachable, error: %q", err)
		}
	})

	testCases := []struct {
		path   string
		method string
		body   io.Reader
		want   string
	}{
		{"/versao", http.MethodGet, nil, "0.0.1"},
	}

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
