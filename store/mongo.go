package store

import (
	"context"
	"log"
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

	for _, linha := range linhas {
		filtro := bson.M{"cod": linha.Codigo}

		var l model.Linha
		err = coll.FindOne(ms.ctx, filtro).Decode(&l)

		switch {
		case err == nil:
			log.Printf("linha %d encontrada. Atualizando data de atualização.", l.Codigo)
			l.AtualizadoEm = time.Now().UnixNano()
		case err.Error() == "mongo: no documents in result":
			log.Printf("linha %d não encontrada. Inserindo linha na base.", l.Codigo)
			l = linha
			l.AtualizadoEm = time.Now().UnixNano()
			l.CriadoEm = time.Now().UnixNano()
		default:
			log.Printf("Erro ao inserir Linhas: %q", err)
			return
		}

		operation := mongo.NewUpdateOneModel()
		operation.SetUpsert(true)
		operation.SetFilter(filtro)
		operation.SetUpdate(l)
		operations = append(operations, operation)
	}

	_, err = coll.BulkWrite(ms.ctx, operations)

	return
}

// Disconnect desconecta a store do banco
func (ms *MongoStore) Disconnect() {
	ms.client.Disconnect(ms.ctx)
}
