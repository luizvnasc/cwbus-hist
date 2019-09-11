package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {

	testCases := []struct {
		nome    string
		path    string
		method  string
		handler HandlerFunc
		want    string
	}{
		{
			nome:    "versao",
			path:    "/versao",
			method:  http.MethodGet,
			handler: versao,
			want:    "0.0.1",
		},
		{
			nome:    "notFound",
			path:    "/",
			method:  http.MethodGet,
			handler: notFound,
			want:    "Rota não encontrada.",
		},
		{
			nome:    "notFound",
			path:    "/teste",
			method:  http.MethodPost,
			handler: notFound,
			want:    "Rota não encontrada.",
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testando handler %q", tc.nome), func(t *testing.T) {
			req, _ := http.NewRequest(tc.path, tc.method, nil)
			res := httptest.NewRecorder()

			http.HandlerFunc(tc.handler).ServeHTTP(res, req)

			got := res.Body.String()

			if got != tc.want {
				t.Errorf("Erro ao testar handler de rota. Retorno esperado: %q, retorno obtido: %q", tc.want, got)
			}
		})
	}

}
