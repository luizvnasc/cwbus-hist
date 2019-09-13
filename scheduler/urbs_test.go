package scheduler

import (
	"testing"

	"github.com/robfig/cron/v3"
)

func TestUrbsScheduler(t *testing.T) {
	c := cron.New()
	client := mongp
	s := store.NewMongoStore
	t.Run("Criar Urbs Scheduler", func(t *testing.T) {
		
		scheduler, err := NewUrbsScheduler(c,)
	})
}
