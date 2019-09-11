//Package router contém as rotas do sistema.
package router

import "net/http"

// HandlerFunc é a interface de uma função handler http.
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Route representa uma rota no sistema.
type Route struct {
	path    string
	methods []string
	handler HandlerFunc
}

// Path da rota
func (r *Route) Path() string {
	return r.path
}

// Handler da rota
func (r *Route) Handler() HandlerFunc {
	return r.handler
}

// Methods HTTP aceitos pela rota.
func (r *Route) Methods() []string {
	return r.methods
}

// NewRoute cria uma rota.
func NewRoute(path string, methods []string, handler HandlerFunc) *Route {
	return &Route{path, methods, handler}
}

// Routes é um mapa das rotas do sistema.
func Routes() []*Route {
	return []*Route{
		NewRoute("/versao", []string{http.MethodGet}, versao),
		NewRoute("/", []string{http.MethodGet}, notFound),
	}
}
