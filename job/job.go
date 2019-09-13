// Package job contém os jobs que serão executados para consumir os serviços da urbs
package job

// Job é um trabalho que será executado de acordo com sua especificação.
type Job struct {
	spec string
	task func()
}

// Spec getter
func (j *Job) Spec() string {
	return j.spec
}

// New é um construtor de Job
func New(spec string, task func()) *Job {
	return &Job{spec,task}
}
