package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestWakeUp(t *testing.T) {
	counter := 0
	want := "Trabalho... Trabalho..."
	spy := func(r *http.Response) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Erro ao ler body da requisição")
		}
		defer r.Body.Close()

		if string(b) != want {
			t.Fatalf("Resposta diferente da espertada. Obtere %q, esperava %q", string(b), want)
		}
		counter++
	}
	sleepTime := 5 * time.Second

	wakeUp("SUA URL AQUI", sleepTime, spy)
}
