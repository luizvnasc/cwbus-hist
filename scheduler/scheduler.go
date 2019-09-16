// Package jobs contém os jobs que serão executados para consumir os serviços da urbs
package scheduler

// Error é um erro do pacote scheduler
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrNoCron : "Cron não informado"
	ErrNoCron = Error("Cron não informado")
)

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
func NewJob(spec string, task func()) *Job {
	return &Job{spec, task}
}

//Scheduler é uma interface para os agendadores de tarefas
type Scheduler interface {
	Execute()
	Terminate()
}
