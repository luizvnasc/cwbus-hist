package scheduler

import (
	"context"
	"os"
	"testing"

	"github.com/luizvnasc/cwbus-hist/db"
	"github.com/luizvnasc/cwbus-hist/store"
	"github.com/robfig/cron/v3"
)

func TestUrbsScheduler(t *testing.T) {
	c := cron.New()
	ctx := context.Background()
	client, err := db.NewMongoClient(ctx, os.Getenv("CWBUS_DB_URL"))
	if err != nil {
		t.Fatalf("Erro ao criar client mongo: %v", err)
	}
	s := store.NewMongoStore(ctx, client)
	t.Run("Criar Urbs Scheduler", func(t *testing.T) {

		scheduler, err := NewUrbsScheduler(c, s)

		if err != nil {
			t.Errorf("Erro ao criar scheduler de jobs da urbs: %v", err)
		}
		if scheduler == nil {
			t.Error("Erro ao criar scheduler de jobs da urbs: scheduler é nulo")
		}
	})
	t.Run("Criar Urbs Scheduler sem informar cron", func(t *testing.T) {

		_, err := NewUrbsScheduler(nil, s)

		if err == nil {
			t.Errorf("Erro ao criar scheduler de jobs da urbs: Esperava-se %q, obteve-se %v", ErrNoCron, err)
		}

	})
}
