package scheduler

import (
	"github.com/luizvnasc/cwbus-hist/store"
	"github.com/robfig/cron/v3"
)

// UrbsScheduler é um agenda os jobs referentes ao serviço da urbs que serão executados.
type UrbsScheduler struct {
	cron  *cron.Cron
	store store.Storer
}

// NewUrbsScheduler é um construtor da estrutura UrbsScheduler
func NewUrbsScheduler(c *cron.Cron, store store.Storer) (*UrbsScheduler, error) {
	if c == nil {
		return nil, ErrNoCron
	}
	return &UrbsScheduler{c, store}, nil
}
