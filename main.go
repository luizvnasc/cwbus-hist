// cwbus-hist é uma aplicação que monitora os serviços de dados abertos da URBS sobre o transporte
// público de Curitiba e armazena um histórico em uma base pública.
package main

import (
	"os"

	"github.com/luizvnasc/cwbus-hist/jobs"
	"github.com/luizvnasc/cwbus-hist/server"
)

func main() {
	// ctx := context.Background()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("CWBUS_DB_URL")))
	// if err != nil {
	// 	log.Fatalf("Erro ao conectar no banco: %q", err)
	// 	os.Exit(1)
	// }

	//store := store.NewMongoStore(ctx, client)
	jobs.Execute()
	app := server.New(os.Getenv("PORT"))
	app.Run()
	//store.Disconnect()
}
