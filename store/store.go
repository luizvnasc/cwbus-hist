// Package store contém os métodos que irão armazenar os dados consumidos em uma base
package store

import "github.com/luizvnasc/cwbus.io/model"

//Storer é a representação de como será implementada a Store
type Storer interface {
	SaveLinhas(linhas model.Linhas) error
}
