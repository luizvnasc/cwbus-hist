package test

import "github.com/luizvnasc/cwbus-hist/model"

// MockStore é uma store para ser usada em testes
type MockStore struct {
	linhas   model.Linhas
	veiculos model.Veiculos
}

// SaveLinhas salva as linhas na store
func (ms *MockStore) SaveLinhas(linhas model.Linhas) error {
	ms.linhas = linhas
	return nil
}

// Linhas retorna as linhas salvas na store
func (ms *MockStore) Linhas() (model.Linhas, error) {
	return ms.linhas, nil
}

// SaveVeiculos salva os veículos na store.
func (ms *MockStore) SaveVeiculos(veiculos map[string]model.Veiculo) (err error) {
	ms.veiculos = model.Veiculos{}
	for _, veiculo := range veiculos {
		ms.veiculos = append(ms.veiculos, veiculo)
	}
	return
}

// Veiculos retorna os veículos salvos na store
func (ms *MockStore) Veiculos() (model.Veiculos, error) {
	return ms.veiculos, nil
}
