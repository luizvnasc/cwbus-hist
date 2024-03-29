// cwbus-hist é uma aplicação que monitora os serviços de dados abertos da URBS sobre o transporte
// público de Curitiba e armazena um histórico em uma base pública.
package main

import (
	"context"
	"log"
	"os"

	"github.com/luizvnasc/cwbus-hist/config"
	"github.com/luizvnasc/cwbus-hist/db"
	"github.com/luizvnasc/cwbus-hist/scheduler"
	"github.com/luizvnasc/cwbus-hist/server"
	"github.com/luizvnasc/cwbus-hist/store"
)

func main() {
	config := &config.EnvConfigurer{}

	log.Println("Criando cliente mongodb")
	ctx := context.Background()
	client, err := db.NewMongoClient(ctx, config.DBStrConn())
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %q", err)
		os.Exit(1)
	}

	log.Println("Criando store")
	s := store.NewMongoStore(ctx, client, config)

	log.Println("Iniciando Schedulers")
	appScheduler := scheduler.NewAppScheduler(config)
	urbsScheduler, err := scheduler.NewUrbsScheduler(s, config)
	if err != nil {
		log.Fatalf("Erro ao iniciar o schduler da urbs")
	}

	schedulers := []scheduler.Scheduler{appScheduler, urbsScheduler}

	for _, s := range schedulers {
		s.Execute()
		defer s.Terminate()
	}

	log.Println("Iniciando servidor")
	app := server.New(os.Getenv("PORT"))
	app.Run()

}
