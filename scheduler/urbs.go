package scheduler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/luizvnasc/cwbus-hist/model"
	"github.com/luizvnasc/cwbus-hist/store"
	"github.com/robfig/cron/v3"
)

// UrbsScheduler é um agenda os jobs referentes ao serviço da urbs que serão executados.
type UrbsScheduler struct {
	cron  *cron.Cron
	store store.Storer
	code  string
	jobs  Jobs
}

// Task que recupera as linhas do serviço da urbs e salva no banco.
func (us *UrbsScheduler) getLinhas() {
	res, err := http.Get(fmt.Sprintf("http://transporteservico.urbs.curitiba.pr.gov.br/getLinhas.php?c=%s", us.code))
	if err != nil {
		log.Printf("Erro ao obter Linhas: %q", err)
		return
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Erro ao let body do serviço getLinhas: %q", err)
		return
	}
	defer res.Body.Close()

	var linhas model.Linhas
	if err := json.Unmarshal(result, &linhas); err != nil {
		log.Printf("Erro ao converter json de linhas para struct Linha: %q", err)
		return
	}

	if err := us.store.SaveLinhas(linhas); err != nil {
		log.Printf("Erro ao salvar linhas no banco: %q", err)
		return
	}

}

func (us *UrbsScheduler) getPontosLinhas(codigo string) (pontos model.Pontos, err error) {
	res, err := http.Get(fmt.Sprintf("http://transporteservico.urbs.curitiba.pr.gov.br/getPontosLinha.php?linha=%s&c=%s", codigo, us.code))
	if err != nil {
		log.Printf("Erro ao obter Linhas: %q", err)
		return
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Erro ao let body do serviço getLinhas: %q", err)
		return
	}
	defer res.Body.Close()

	if err = json.Unmarshal(result, &pontos); err != nil {
		log.Printf("Erro ao converter json de linhas para struct Linha: %q", err)
		return
	}
	return
}

// Execute inicia a execução dos jobs do scheduler
func (us *UrbsScheduler) Execute() {
	for _, job := range us.jobs {
		us.cron.AddFunc(job.Spec(), job.Task())
	}
	us.cron.Start()
}

// Terminate para a execução do scheduler
func (us *UrbsScheduler) Terminate() {
	us.cron.Stop()
}

// NewUrbsScheduler é um construtor da estrutura UrbsScheduler
func NewUrbsScheduler(store store.Storer) (*UrbsScheduler, error) {
	c := cron.New()
	code := os.Getenv("CWBUS_URBS_CODE")
	if len(code) == 0 {
		return nil, ErrNoUrbsCode
	}
	scheduler := &UrbsScheduler{cron: c, store: store, code: code}
	scheduler.jobs = append(scheduler.jobs, NewJob("@hourly", scheduler.getLinhas))
	return scheduler, nil
}
