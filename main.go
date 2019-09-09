// cwbus.io é uma aplicação que monitora os serviços de dados abertos da URBS sobre o transporte 
// público de Curitiba e armazena um histórico em uma base pública.
package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("CWBUS_DB_URL")))
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %q", err)
		os.Exit(1)
	}

	store := store.NewMongoStore(ctx, client)
	//TODO: CONSUMERS VEM AQUI
	store.Disconnect()
}
