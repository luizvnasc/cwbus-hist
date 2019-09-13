// Package task contém as tarefas que serão executadas pelos jobs do sistema.
// Criei um package diferente para deixar legível.
package task

import (
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// WakeUp acorda o dyno do heroku
func WakeUp() {
	url := os.Getenv("CWBUS_WAKEUP_URL")
	res, err := http.Get(url)
	if err != nil {
		errors.Errorf("Erro ao acordar o dyno: %q", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		errors.Errorf("Erro ao acordar o dyno, Status %q", res.StatusCode)
	}
	log.Println("Trabalho... Trabalho...")
}
