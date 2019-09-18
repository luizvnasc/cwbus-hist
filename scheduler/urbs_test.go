package scheduler

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/luizvnasc/cwbus-hist/db"
	"github.com/luizvnasc/cwbus-hist/model"
	"github.com/luizvnasc/cwbus-hist/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestUrbsScheduler(t *testing.T) {
	ctx := context.Background()
	client, err := db.NewMongoClient(ctx, os.Getenv("CWBUS_DB_URL"))
	if err != nil {
		t.Fatalf("Erro ao criar client mongo: %v", err)
	}
	s := store.NewMongoStore(ctx, client)

	t.Run("Criar Urbs Scheduler", func(t *testing.T) {

		scheduler, err := NewUrbsScheduler(s)

		if err != nil {
			t.Errorf("Erro ao criar scheduler de jobs da urbs: %v", err)
		}
		if scheduler == nil {
			t.Error("Erro ao criar scheduler de jobs da urbs: scheduler é nulo")
		}
	})

	t.Run("getLinhas Task Caminho feliz", func(t *testing.T) {
		scheduler, _ := NewUrbsScheduler(s)
		scheduler.getLinhas()
		linhas := client.Database(os.Getenv("CWBUS_DB_HIST")).Collection("linhas")
		AssertNumberOfDocuments(ctx, t, linhas, 311)
	})

	t.Run("getPontos Task Caminho feliz", func(t *testing.T) {
		scheduler, _ := NewUrbsScheduler(s)

		var wg sync.WaitGroup

		errChan := make(chan error, 1)
		dataChan := make(chan model.Pontos, 1)
		defer close(errChan)
		defer close(dataChan)
		wg.Add(1)
		go scheduler.getPontos(&wg, errChan, dataChan, "464")
		wg.Wait()

		select {
		case <-errChan:
			t.Errorf("Erro ao obter pontos de uma linha: %q", err)
		case pontos := <-dataChan:
			if len(pontos) != 59 {
				t.Errorf("Erro ao contar os pontos da linha. Esperava-se %d, obteve-se %d", 59, len(pontos))
			}
		}

	})

	t.Run("getPontosLinhas", func(t *testing.T) {
		scheduler, _ := NewUrbsScheduler(s)
		linhas, _ := scheduler.store.Linhas()
		//Reinicia os pontos das linhas
		for i := range linhas {
			linhas[i].Pontos = model.Pontos{}
		}
		linhas, err := scheduler.getPontosLinhas(linhas)
		if err != nil {
			t.Errorf("Erro ao obter os pontos das linhas: %q", err)
		}
		for i := range linhas {
			if len(linhas[i].Pontos) == 0 {
				t.Errorf("Erro ao obter os pontos das linhas %q. Pontos: %v", linhas[i].Codigo, linhas[i].Pontos)
			}
		}
	})

	t.Run("Criar Urbs Scheduler sem informar sem código urbs", func(t *testing.T) {
		os.Setenv("CWBUS_URBS_CODE", "")
		_, err := NewUrbsScheduler(s)

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
