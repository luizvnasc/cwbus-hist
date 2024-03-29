package store

import (
	"context"
	"time"

	"github.com/luizvnasc/cwbus-hist/config"
	"github.com/luizvnasc/cwbus-hist/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoStore é uma store que se comunica com uma base de dados mongodb.
type MongoStore struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

// NewMongoStore cria uma Store para uma base de dados mongodb.
func NewMongoStore(ctx context.Context, client *mongo.Client, config config.Configurer) (store Storer) {
	store = &MongoStore{client: client, db: client.Database(config.DBName())}
	return
}

// SaveLinhas salva uma lista de linhas na base de dados.
func (ms *MongoStore) SaveLinhas(linhas model.Linhas) (err error) {
	var operations []mongo.WriteModel
	coll := ms.db.Collection("linhas")
	filtro := bson.M{}

	var linhasCadastradas model.Linhas
	cur, err := coll.Find(ms.ctx, filtro)

	if err != nil {
		return err
	}
	for cur.Next(ms.ctx) {
		var linha model.Linha
		err := cur.Decode(&linha)
		if err != nil {
			return err
		}
		linhasCadastradas = append(linhasCadastradas, linha)
	}

	for _, linha := range linhas {
		var l *model.Linha
		for _, linhaCadastrada := range linhasCadastradas {
			if linhaCadastrada.Codigo == linha.Codigo {
				l = &linhaCadastrada
				l.Pontos = linha.Pontos
				l.Tabela = linha.Tabela
				break
			}
		}
		if l != nil {
			//fmt.Printf("linha %q já cadastrada\n", l.Codigo)
			l.AtualizadoEm = time.Now().Unix()
		} else {
			//fmt.Printf("linha %q é nova\n", l.Codigo)
			l = &linha
			l.AtualizadoEm = time.Now().Unix()
			l.CriadoEm = time.Now().Unix()
		}

		operation := mongo.NewUpdateOneModel()
		operation.SetUpsert(true)
		operation.SetFilter(bson.M{"cod": l.Codigo})
		operation.SetUpdate(*l)
		operations = append(operations, operation)
	}

	_, err = coll.BulkWrite(ms.ctx, operations)

	return
}

// Linhas lista as linhas armazenadas no banco.
func (ms *MongoStore) Linhas() (linhas model.Linhas, err error) {
	cur, err := ms.db.Collection("linhas").Find(ms.ctx, bson.D{})
	if err != nil {
		return
	}
	for cur.Next(ms.ctx) {
		var linha model.Linha
		err = cur.Decode(&linha)
		if err != nil {
			linhas = model.Linhas{}
			return
		}
		linhas = append(linhas, linha)
	}
	return
}

// SaveVeiculos carrega a coleção veiculos com uma lista de veiculos.
func (ms *MongoStore) SaveVeiculos(veiculos map[string]model.Veiculo) error {
	var operations []mongo.WriteModel
	coll := ms.db.Collection("veiculos")

	// Limpa a base para a adicionar as novas situações dos veículos.
	//coll.DeleteMany(ms.ctx, bson.M{})

	for _, veiculo := range veiculos {
		veiculo.CriadoEm = time.Now().UTC()
		operation := mongo.NewInsertOneModel()
		operation.SetDocument(veiculo)
		operations = append(operations, operation)
	}

	_, err := coll.BulkWrite(ms.ctx, operations)

	return err
}

// Veiculos lista os veiculos da coleção veiculos.
func (ms *MongoStore) Veiculos() (veiculos model.Veiculos, err error) {
	cur, err := ms.db.Collection("veiculos").Find(ms.ctx, bson.D{})
	if err != nil {
		return
	}
	for cur.Next(ms.ctx) {
		var veiculo model.Veiculo
		err = cur.Decode(&veiculo)
		if err != nil {
			veiculos = []model.Veiculo{}
			return
		}
		veiculos = append(veiculos, veiculo)
	}
	return
}

// Disconnect desconecta a store do banco
func (ms *MongoStore) Disconnect() {
	ms.client.Disconnect(ms.ctx)
}
