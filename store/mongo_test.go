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

func TestLinhas(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createMongoClient(ctx, t)
	store := NewMongoStore(ctx, client)

	var linhas = model.Linhas{
		model.Linha{
			Codigo:           "464",
			Nome:             "A. MUNHOZ / J. BOTÂNICO",
			SomenteCartao:    "S",
			CategoriaServico: "CONVENCIONAL",
			Cor:              "AMARELA",
		},
	}

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

func TestVeiculos(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createMongoClient(ctx, t)
	store := NewMongoStore(ctx, client)

	veiculos := model.Veiculos{
		"GC295": model.Veiculo{
			Codigo:      "GC295",
			Refresh:     "15:05",
			Latitude:    "-25.443406",
			Longitude:   "-49.270213",
			CodigoLinha: "666",
			Adaptado:    "1",
			Tipo:        "1",
			Tabela:      "2",
			Situacao1:   "ATRASADO",
			Situacao2:   "TIPO INCOMPATIVEL",
			Sent:        "VOLTA",
			Tcount:      1,
			Sentido:     "198-BAIRRO NOVO MUNDO (15:38)",
		},
	}

	t.Run("Inserindo veiculos no banco", func(t *testing.T) {
		err := store.SaveVeiculos(veiculos)
		if err != nil {
			t.Errorf("Erro ao salvar Veiculos no BD: %q", err)
		}
	})

	t.Run("Listar Veiculos do banco", func(t *testing.T) {
		veiculos, err := store.Veiculos()
		if err != nil {
			t.Errorf("Erro ao obter os veiculos cadastradas: %q", err)
		}
		if len(veiculos) != 311 {
			t.Errorf("Erro ao contar Veiculos. Esperava-se %d veiculos, obteve-se %d", 1, len(veiculos))
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
