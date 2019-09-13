package scheduler

import "testing"

func TestJob(t *testing.T) {
	t.Run("Criando Job", func(t *testing.T) {
		cronSpec := "* * * * * ? *"
		task := func() { println("Olá job") }
		j := New(cronSpec, task)
		if j == nil {
			t.Errorf("Erro ao criar job, esperava-se um objeto, obteve-se nil")
		}
		if j.Spec() != cronSpec {
			t.Errorf("Erro na criação do job, experava-se especificação %q, obteve-se %q", cronSpec, j.Spec())
		}
	})
}
