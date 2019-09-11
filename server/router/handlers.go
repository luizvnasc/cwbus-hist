package router

import (
	"fmt"
	"net/http"
)

// Versao é retorna a versao do sistema.
func versao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "0.0.1")
}
