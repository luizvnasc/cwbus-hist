package test

import "testing"

// AssertStringsEqual verifica se duas string são iguais
func AssertStringsEqual(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("Esperava-se %q, obteve-se %q", want, got)
	}
}
