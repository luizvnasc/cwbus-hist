package scheduler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/luizvnasc/cwbus-hist/model"
	"github.com/luizvnasc/cwbus-hist/store"
	"github.com/robfig/cron/v3"
)

// UrbsScheduler é um agenda os jobs referentes ao serviço da urbs que serão executados.
type UrbsScheduler struct {
	cron       *cron.Cron
	store      store.Storer
	code       string
	serviceURL string
	jobs       Jobs
}

// Task que recupera as linhas do serviço da urbs e salva no banco.
func (us *UrbsScheduler) getLinhas() {
	log.Println("Obtendo linhas...")
	res, err := http.Get(fmt.Sprintf("%s/getLinhas.php?c=%s", us.serviceURL, us.code))
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

	us.getPontosLinhas(linhas)

	if err := us.store.SaveLinhas(linhas); err != nil {
		log.Printf("Erro ao salvar linhas no banco: %q", err)
		return
	}
	log.Println("Linhas obtidas com sucesso.")
}

// getPontosLinhas recebe como parâmetro uma lista de linhas e armazena seus respectivos pontos.
func (us *UrbsScheduler) getPontosLinhas(linhas model.Linhas) (model.Linhas, error) {
	errChannels := make([]chan error, len(linhas))
	dataChannels := make([]chan model.Pontos, len(linhas))

	for i := range errChannels {
		errChannels[i] = make(chan error, 1)
		dataChannels[i] = make(chan model.Pontos, 1)
		defer close(errChannels[i])
		defer close(dataChannels[i])
	}

	var wg sync.WaitGroup
	wg.Add(len(linhas))

	for i, linha := range linhas {
		go us.getPontos(&wg, errChannels[i], dataChannels[i], linha.Codigo)
		time.Sleep(3 * time.Millisecond) //evita reset de conexão
	}
	wg.Wait()

	for i := range linhas {
		select {
		case err := <-errChannels[i]:
			log.Printf("Erro ao obter pontos da linha %s: %q", linhas[i].Codigo, err)
			return model.Linhas{}, nil
		case pontos := <-dataChannels[i]:
			linhas[i].Pontos = pontos
		}
	}
	return linhas, nil
}

// getPontos obtém os pontos de uma determinada linha
func (us *UrbsScheduler) getPontos(wg *sync.WaitGroup, errChan chan error, dataChan chan model.Pontos, codigo string) {
	defer wg.Done()
	res, err := http.Get(fmt.Sprintf("%s/getPontosLinha.php?linha=%s&c=%s", us.serviceURL, codigo, us.code))
	if err != nil {
		errChan <- err
		return
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		errChan <- err
		return
	}
	defer res.Body.Close()
	var pontos model.Pontos
	if err = json.Unmarshal(result, &pontos); err != nil {
		errChan <- err
		return
	}
	dataChan <- pontos

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
	code := os.Getenv("CWBUS_URBS_CODE")
	serviceURL := os.Getenv("CWBUS_URBS_SERVICE_URL")
	if len(code) == 0 {
		return nil, ErrNoUrbsCode
	}
	if len(serviceURL) == 0 {
		return nil, ErrNoServiceURL
	}
	scheduler := &UrbsScheduler{cron: cron.New(),
		store:      store,
		code:       code,
		serviceURL: serviceURL,
	}
	scheduler.jobs = append(scheduler.jobs, NewJob("0 5 * * *", scheduler.getLinhas))
	return scheduler, nil
}
