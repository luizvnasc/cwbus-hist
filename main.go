// cwbus.io é uma aplicação que monitora os serviços de dados abertos da URBS sobre o transporte
// público de Curitiba e armazena um histórico em uma base pública.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/luizvnasc/cwbus.io/store"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("CWBUS_DB_URL")))
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %q", err)
		os.Exit(1)
	}

	store := store.NewMongoStore(ctx, client)

	http.HandleFunc("/versao",Versao)
	log.Fatalf("%q\n", http.ListenAndServe(Addr(),nil))

	store.Disconnect()
}

// Versao é handler de chamada http para o caminho /versao.
func Versao(w http.ResponseWriter, r *http.Request ){
	fmt.Fprintf(w,"0.0.1")
}

// Addr cria o endereço onde a aplicação estará rodando.
func Addr() string {
	if len(os.Getenv("$PORT")) == 0 {
		return ":8081"
	}
	return ":" + os.Getenv("$PORT")
}

func wakeUp(url string, sleepTime int64){

}
