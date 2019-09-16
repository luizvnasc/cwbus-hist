// Package db contém os métodos necessários para se conectar a uma base de dados.
package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoClient retorna um cliente conectado a base mongo.
func NewMongoClient(ctx context.Context, connStr string) (*mongo.Client, error) {
	if len(connStr) == 0 {
		return nil, ErrNoConnString
	}
	if ctx == nil {
		return nil, ErrNoContext
	}
	return mongo.Connect(ctx, options.Client().ApplyURI(connStr))
}
