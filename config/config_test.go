package config

import (
	"os"
	"testing"

	"github.com/luizvnasc/cwbus-hist/test"
)

func TestConfig(t *testing.T) {

	t.Run("Obtendo url de serviços da urbs", func(t *testing.T) {
		want := os.Getenv("CWBUS_URBS_SERVICE_URL")
		got := ServiceURL()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo codigo urbs", func(t *testing.T) {
		want := os.Getenv("CWBUS_URBS_CODE")
		got := UrbsCode()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo conexão com o banco", func(t *testing.T) {
		want := os.Getenv("CWBUS_DB_URL")
		got := DBStrConn()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo nome do banco", func(t *testing.T) {
		want := os.Getenv("CWBUS_DB_HIST")
		got := DBName()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo url para acordar dyno", func(t *testing.T) {
		want := os.Getenv("CWBUS_WAKEUP_URL")
		got := WakeUpURL()
		test.AssertStringsEqual(t, want, got)
	})

}
