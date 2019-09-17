package scheduler

import (
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
)

// AppScheduler é um scheduler para jobs da aplicação que não são referentes aos serviços da urbs.
type AppScheduler struct {
	cron *cron.Cron
	jobs Jobs
}

func (as *AppScheduler) wakeUpDyno() {
	url := os.Getenv("CWBUS_WAKEUP_URL")
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao acordar o dyno: %q", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("Erro ao acordar o dyno, Status: %d", res.StatusCode)
		return
	}
	log.Println("Trabalho... Trabalho...")
}

// Execute inicia a execução dos jobs do AppScheduler
func (as *AppScheduler) Execute() {
	for _, job := range as.jobs {
		as.cron.AddFunc(job.Spec(), job.Task())
	}
	as.cron.Start()
}

// Terminate finaliza a execução dos jobs do AppScheduler
func (as *AppScheduler) Terminate() {
	as.cron.Stop()
}

// NewAppScheduler é um construtor de um AppScheduler
func NewAppScheduler() *AppScheduler {
	c := cron.New()
	appScheduler := &AppScheduler{cron: c}
	appScheduler.jobs = append(appScheduler.jobs, NewJob("*/3 * * * *", appScheduler.wakeUpDyno))
	return appScheduler
}
