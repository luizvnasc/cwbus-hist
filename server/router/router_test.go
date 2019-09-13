package router

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func genericHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Não faço nada")
}

func TestRouter(t *testing.T) {
	t.Run("Criando uma rota não informando o path", func(t *testing.T) {

		r := NewRoute("", []string{http.MethodGet}, genericHandler)
		if r != nil {
			t.Errorf("Erro ao criar rota, esperava-se que ela fosse nil, retornou um objeto")
		}
	})
	t.Run("Criando uma rota não informando o método", func(t *testing.T) {

		r := NewRoute("/teste", []string{}, genericHandler)
		if r != nil {
			t.Errorf("Erro ao criar rota, esperava-se que ela fosse nil, retornou um objeto")
		}
	})

	t.Run("Criando uma rota não informando o handler", func(t *testing.T) {

		r := NewRoute("/teste", []string{http.MethodDelete}, nil)
		if r != nil {
			t.Errorf("Erro ao criar rota, esperava-se que ela fosse nil, retornou um objeto")
		}
	})

	t.Run("Criando uma rota", func(t *testing.T) {
		path := "/teste"
		methods := []string{http.MethodDelete}
		r := NewRoute(path, methods, genericHandler)
		if r == nil {
			t.Errorf("Erro ao criar uma rota, esperava-se uma rota, retornou nil")
		}
		if path != r.Path() {
			t.Errorf("Erro ao validar path da rota, esperava-se %q, obteve-se %q", path, r.Path())
		}
		if !reflect.DeepEqual(methods, r.Methods()) {
			t.Errorf("Erro ao validar métodos da rota, esperava-se %v, obteve-se %v", methods, r.Methods())
		}
	})
}
