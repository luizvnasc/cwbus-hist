// Package store contém os métodos que irão armazenar os dados consumidos em uma base
package store

import "github.com/luizvnasc/cwbus-hist/model"

//Storer é a representação de como será implementada a Store
type Storer interface {
	// Salva as linhas no banco de dados
	SaveLinhas(linhas model.Linhas) error
	// Recupera as linhas dos bancos de dados
	Linhas() (model.Linhas, error)
	// Salva os veiculos no banco
	SaveVeiculos(veiculos map[string]model.Veiculo) error
	// Lista os veiculos do banco
	Veiculos() (model.Veiculos, error)
}
