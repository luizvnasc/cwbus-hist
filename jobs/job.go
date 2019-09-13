// Package jobs contém os jobs que serão executados para consumir os serviços da urbs
package jobs

import (
	"log"

	"github.com/luizvnasc/cwbus-hist/jobs/task"
	"github.com/robfig/cron/v3"
)

var c *cron.Cron

// Job é um trabalho que será executado de acordo com sua especificação.
type Job struct {
	spec string
	task func()
}

// Spec getter
func (j *Job) Spec() string {
	return j.spec
}

// Spec getter
func (j *Job) Task() func() {
	return j.task
}

// Jobs é uma lista de jobs
type Jobs []*Job

// New é um construtor de Job
func New(spec string, task func()) *Job {
	return &Job{spec, task}
}

var jobs = Jobs{
	New("*/3 * * * *", task.WakeUp),
}

// Execute inicia a execução dos jobs.
func Execute() {
	c = cron.New()
	for _, job := range jobs {
		id, err := c.AddFunc(job.Spec(), job.Task())
		if err != nil {
			log.Fatalf("Erro ao adicionar job: %q", err)
		}
		log.Printf("Job adicionado. ID: %q",id)
	}
	log.Printf("Iniciando Jobs...")
	c.Start()
	
}

// Terminate para a execução dos jobs.
func Terminate() {
	log.Printf("Finalizando Jobs...")
	c.Stop()
}
