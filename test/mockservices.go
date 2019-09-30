package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
)

func NewMockServer(handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}

// GetVeiculosHandler simula a o serviço getVeiculos da urbs
func GetVeiculosHandler(w http.ResponseWriter, r *http.Request) {
	path, _ := filepath.Abs("../test/getVeiculos.json")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	w.Write([]byte(b))
}

// GetVeiculosWrongBodyHandler simula um retorno errado do serviço getveiculos
func GetVeiculosWrongBodyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("teste"))
}

// GetVeiculosStatus500Handler simula um status 500 do serviços getVeiculos
func GetVeiculosStatus500Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
