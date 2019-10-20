package store

import (
	"context"
	"testing"
	"time"

	"github.com/luizvnasc/cwbus-hist/config"
	"github.com/luizvnasc/cwbus-hist/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestLinhas(t *testing.T) {
	config := &config.EnvConfigurer{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createMongoClient(ctx, t, config)
	store := NewMongoStore(ctx, client, config)

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
	config := &config.EnvConfigurer{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := createMongoClient(ctx, t, config)
	store := NewMongoStore(ctx, client, config)

	veiculos := map[string]model.Veiculo{
		"GC295": model.Veiculo{
			Codigo:      "GC295",
			Refresh:     "15:05",
			Latitude:    "-25.443406",
			Longitude:   "-49.270213",
			CodigoLinha: "666",
			Adaptado:    "1",
			Tipo:        "1",
			Tabela:      "2",
			Situacao:    "ATRASADO",
			Situacao2:   "TIPO INCOMPATIVEL",
			Sent:        "VOLTA",
			Tcount:      1,
			CriadoEm:    time.Now(),
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
		count := 0;
		for _, v := veiculos {
			if v.Codigo == "GC295" {
				count++
			}
		}
		if count > 3 {
			t.Errorf("Erro ao contar Veiculos. Esperava-se no máximo %d veiculos, obteve-se %d", 3, len(veiculos))
		}
	})
}

// Helper que cria uma conexão com a base de dados.
func createMongoClient(ctx context.Context, t *testing.T, config config.Configurer) *mongo.Client {
	t.Helper()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBStrConn()))
	if err != nil {
		t.Fatalf("Erro ao criar conexão com o banco: %q", err)
	}
	return client
}
