package db

import (
	"context"
	"os"
	"testing"
)

func TestMongo(t *testing.T) {
	ctx := context.Background()
	t.Run("Cria uma instância do cliente mongo sem informar a string de conexão", func(*testing.T) {
		_, err := NewMongoClient(ctx, "")
		if err != ErrNoConnString {
			t.Errorf("Erro ao criar cliente mongo. Esperava-se %q, obteve-se %q", ErrNoConnString, err)
		}
	})
	t.Run("Cria uma instância do cliente mongo sem informar a contexto", func(*testing.T) {
		_, err := NewMongoClient(nil, "teste")
		if err != ErrNoContext {
			t.Errorf("Erro ao criar cliente mongo. Esperava-se %q, obteve-se %q", ErrNoContext, err)
		}
	})
	t.Run("Cria uma instância do cliente mongo corretamente", func(*testing.T) {
		client, err := NewMongoClient(ctx, os.Getenv("CWBUS_DB_URL"))
		if err != nil {
			t.Errorf("Erro ao criar cliente mongo. Esperava-se nil, obteve-se %q", err)
		}

		if client == nil {
			t.Errorf("Cliente nulo")
		}
	})
}
