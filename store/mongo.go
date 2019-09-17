package store

import (
	"context"
	"os"
	"time"

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
func NewMongoStore(ctx context.Context, client *mongo.Client) (store *MongoStore) {
	store = &MongoStore{client: client}
	store.db = store.client.Database(os.Getenv("CWBUS_DB_HIST"))
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

// Disconnect desconecta a store do banco
func (ms *MongoStore) Disconnect() {
	ms.client.Disconnect(ms.ctx)
}
