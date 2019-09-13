package db

import (
	"context"
	"testing"
)

func TestMongo(t *testing.T) {
	ctx := context.Background()
	t.Run("Cria uma instância do cliente mongo sem informar a string de conexão", func(*testing.T) {
		_, err := NewMongoClient(ctx, "")
		if err == nil {
			t.Errorf("Erro ao criar cliente mongo. Esperava-se %q, obteve-se %q", ErrNoConnString, err)
		}
	})
}
