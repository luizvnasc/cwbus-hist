package scheduler

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/luizvnasc/cwbus-hist/config"
	"github.com/luizvnasc/cwbus-hist/test/mock"
)

func TestAppScheduler(t *testing.T) {
	config := &config.EnvConfigurer{}
	mockConfig := &mock.MockConfigurer{}
	t.Run("Cria um scheduler da aplicação", func(t *testing.T) {
		s := NewAppScheduler(config)
		if s == nil {
			t.Errorf("Scheduler não foi criado.")
		}
	})

	cases := []struct {
		status int
		want   string
	}{
		{0, "Erro ao acordar o dyno:"},
		{http.StatusBadRequest, "Erro ao acordar o dyno, Status:"},
		{http.StatusOK, "Trabalho..."},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("Teste ao acordar dyno retornando status: %d", test.status), func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
			}))
			// Close the server when test finishes
			defer server.Close()

			var buf bytes.Buffer
			log.SetOutput(&buf)
			defer func() {
				log.SetOutput(os.Stderr)
			}()

			mockConfig.SetWakeUpURL(server.URL)

			s := NewAppScheduler(mockConfig)
			s.wakeUpDyno()
			got := buf.String()
			if !strings.Contains(got, test.want) {
				t.Errorf("Erro ao validar url na task wakeup. Log deveria conter %q mas retornou %q", test.want, got)
			}
		})
	}
}
