package store

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/luizvnasc/cwbus-hist/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var linhas = model.Linhas{
	model.Linha{
		Codigo:           "464",
		Nome:             "A. MUNHOZ / J. BOTÂNICO",
		SomenteCartao:    "S",
		CategoriaServico: "CONVENCIONAL",
		Cor:              "AMARELA",
	},
}

func TestMongoStore(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createMongoClient(ctx, t)
	store := NewMongoStore(ctx, client)

	t.Run("Inserindo linhas no banco", func(t *testing.T) {
		err := store.SaveLinhas(linhas)
		if err != nil {
			t.Errorf("Erro ao salvar linhas no BD: %q", err)
		}
	})
	t.Run("Listar Linhas do banco", func(t *testing.T) {
		linhas, err := store.Linhas()
		if err != nil {
			t.Errorf("Erro ao obter as linhas cadastradas: %q", err)
		}
		if len(linhas) != 311 {
			t.Errorf("Erro ao contar Linhas. Esperava-se %d linhas, obteve-se %d", 311, len(linhas))
		}
	})
}

// Helper que cria uma conexão com a base de dados.
func createMongoClient(ctx context.Context, t *testing.T) *mongo.Client {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("CWBUS_DB_URL")))
	if err != nil {
		t.Fatalf("Erro ao criar conexão com o banco: %q", err)
	}
	return client
}
