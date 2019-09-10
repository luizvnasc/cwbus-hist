// cwbus.io é uma aplicação que monitora os serviços de dados abertos da URBS sobre o transporte
// público de Curitiba e armazena um histórico em uma base pública.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jasonlvhit/gocron"
)

func main() {
	// ctx := context.Background()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("CWBUS_DB_URL")))
	// if err != nil {
	// 	log.Fatalf("Erro ao conectar no banco: %q", err)
	// 	os.Exit(1)
	// }

	//store := store.NewMongoStore(ctx, client)

	go func() {
		log.Println("Iniciando gocron...")
		gocron.Every(1).Minute().Do(wakeUp)
		<-gocron.Start()
	}()
	log.Println("Iniciando servidor...")
	http.HandleFunc("/versao", Versao)
	log.Fatalf("%q\n", http.ListenAndServe(Addr(), nil))
	//store.Disconnect()
}

// Versao é handler de chamada http para o caminho /versao.
func Versao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "0.0.1")
}

// Addr cria o endereço onde a aplicação estará rodando.
func Addr() string {
	if len(os.Getenv("PORT")) == 0 {
		return ":8081"
	}
	return ":" + os.Getenv("PORT")
}

func wakeUp() {
	url := os.Getenv("CWBUS_VERSAO_URL")
	if len(url) == 0 {
		url = "http://localhost:8081/versao"
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao acodar o dinossáuro: %q\n", err)
	}
	if res.StatusCode == 200 {
		log.Println("Trabalho... Trabalho...")
	} else {
		log.Fatalf("Status retornado diferente do esperado: %d", res.StatusCode)
	}

}
