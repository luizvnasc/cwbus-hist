package config

import (
	"os"
	"testing"

	"github.com/luizvnasc/cwbus-hist/test"
)

func TestConfig(t *testing.T) {
	ec := EnvConfigurer{}
	t.Run("Obtendo url de serviços da urbs", func(t *testing.T) {
		want := os.Getenv("CWBUS_URBS_SERVICE_URL")
		got := ec.ServiceURL()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo codigo urbs", func(t *testing.T) {
		want := os.Getenv("CWBUS_URBS_CODE")
		got := ec.UrbsCode()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo conexão com o banco", func(t *testing.T) {
		want := os.Getenv("CWBUS_DB_URL")
		got := ec.DBStrConn()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo nome do banco", func(t *testing.T) {
		want := os.Getenv("CWBUS_DB_HIST")
		got := ec.DBName()
		test.AssertStringsEqual(t, want, got)
	})

	t.Run("Obtendo url para acordar dyno", func(t *testing.T) {
		want := os.Getenv("CWBUS_WAKEUP_URL")
		got := ec.WakeUpURL()
		test.AssertStringsEqual(t, want, got)
	})

}
