package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

const port = "8081"

func TestServer(t *testing.T) {

	t.Run("Iniciando servidor sem informar a porta", func(t *testing.T) {
		quit := make(chan bool)
		defer close(quit)
		server := runServer(t,"",quit)
		defer server.Shutdown()

		_, err := net.Dial("tcp", ":"+port)
		if err != nil {
			t.Errorf("Site unreachable, error: %q", err)
		}
		quit <- true
	})

	t.Run(fmt.Sprintf("Inciando servidor na porta %q", port), func(t *testing.T) {
		quit := make(chan bool)
		defer close(quit)
		server := runServer(t,port,quit)
		defer server.Shutdown()
		_, err := net.Dial("tcp", ":"+port)
		if err != nil {
			t.Errorf("Site unreachable, error: %q", err)
		}
		quit <- true

	})
}

func TestRoutes(t *testing.T){
	// Testando rotas do servidor
	testCases := []struct {
		path   string
		method string
		body   io.Reader
		want   string
	}{
		{"http://localhost:8081/versao", http.MethodGet, nil, "0.0.1"},
		{"http://localhost:8081/teste", http.MethodGet, nil, "Rota não encontrada."},
	}

	quit := make(chan bool)
	defer close(quit)
	server := runServer(t,port, quit)
	defer server.Shutdown()

	client := &http.Client{}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Testando rota %q", tc.path), func(t *testing.T) {
			req, _ := http.NewRequest(tc.method, tc.path, tc.body)
			res, _ := client.Do(req)
			got, _ := ioutil.ReadAll(res.Body)
			if string(got) != tc.want {
				t.Errorf("Erro ao testar rota %q. Resposta esperada: %q, obtida: %q", tc.path, tc.want, string(got))
			}
		})
	}
	quit <- true
}

func runServer(t *testing.T, port string, quit chan bool) (server *CwbusServer) {
	server = New(port)
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
	return
}
