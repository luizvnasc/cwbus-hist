package scheduler

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/luizvnasc/cwbus-hist/config"
)

func TestAppScheduler(t *testing.T) {
	t.Run("Cria um scheduler da aplicação", func(t *testing.T) {
		s := NewAppScheduler(config.WakeUpURL())
		if s == nil {
			t.Errorf("Scheduler não foi criado.")
		}
	})

	t.Run("Teste acordar dyno com erro de url", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()
		s := NewAppScheduler("teste")
		s.wakeUpDyno()
		got := buf.String()
		if !strings.Contains(got, "Erro ao acordar o dyno:") {
			t.Errorf("Erro ao validar url na task wakeup.")
		}
	})

	t.Run("Teste acordar dyno com erro de statuscode", func(t *testing.T) {
		var buf bytes.Buffer

		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		s := NewAppScheduler("https://httpstat.us/400")
		s.wakeUpDyno()
		got := buf.String()
		if !strings.Contains(got, "Erro ao acordar o dyno, Status:") {
			t.Errorf("Erro ao validar status na task wakeup.")
		}
	})

	t.Run("Teste acordar dyno", func(t *testing.T) {
		var buf bytes.Buffer

		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()
		s := NewAppScheduler("https://httpstat.us/200")
		s.wakeUpDyno()
		got := buf.String()
		if !strings.Contains(got, "Trabalho...") {
			t.Errorf("Erro ao validar status na task wakeup.")
		}
	})
}
