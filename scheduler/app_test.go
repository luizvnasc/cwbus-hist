package scheduler

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestAppScheduler(t *testing.T) {
	t.Run("Cria um scheduler da aplicação", func(t *testing.T) {
		s := NewAppScheduler()
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
		s := NewAppScheduler()
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

		os.Setenv("CWBUS_WAKEUP_URL", "https://httpstat.us/400")

		s := NewAppScheduler()
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

		os.Setenv("CWBUS_WAKEUP_URL", "https://httpstat.us/200")

		s := NewAppScheduler()
		s.wakeUpDyno()
		got := buf.String()
		if !strings.Contains(got, "Trabalho...") {
			t.Errorf("Erro ao validar status na task wakeup.")
		}
	})
}
