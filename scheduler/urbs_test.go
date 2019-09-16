package scheduler

import (
	"context"
	"os"
	"testing"

	"github.com/luizvnasc/cwbus-hist/db"
	"github.com/luizvnasc/cwbus-hist/store"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

		if err != ErrNoCron {
			t.Errorf("Erro ao criar scheduler de jobs da urbs: Esperava-se %q, obteve-se %v", ErrNoCron, err)
		}

	})

	t.Run("getLinhas Task", func(t *testing.T) {
		scheduler, _ := NewUrbsScheduler(c, s)
		scheduler.getLinhas()
		linhas := client.Database(os.Getenv("CWBUS_DB_HIST")).Collection("linhas")
		AssertNumberOfDocuments(ctx, t, linhas, 311)
	})

	t.Run("Criar Urbs Scheduler sem informar sem código urbs", func(t *testing.T) {
		os.Setenv("CWBUS_URBS_CODE", "")
		_, err := NewUrbsScheduler(c, s)

		if err != ErrNoUrbsCode {
			t.Errorf("Erro ao criar scheduler de jobs da urbs: Esperava-se %q, obteve-se %v", ErrNoUrbsCode, err)
		}

	})
}

func AssertNumberOfDocuments(ctx context.Context, t *testing.T, coll *mongo.Collection, want int64) {
	t.Helper()
	count, err := coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		t.Errorf("Erro ao contar os documentos")
	}
	if count != want {
		t.Errorf("Número de documentos errado. Esperado %d, obtido %d", want, count)
	}
}