// Package consumer contém o tipo consumer que será quem irá consumir os serviços da urbs
package consumer

import (
	"fmt"
	"net/http"
)

//Consumer é um tipo que irá consumir um serviço passado a ele.
type Consumer struct {
	url     string
	handler func(res *http.Response) error
	cherr   chan error
}

// Run inicia a execução do consumidor
func (c *Consumer) Run() {
	fmt.Println(c.url)
	res, err := http.Get(c.url)
	if err != nil {
		c.cherr <- err
		return
	}
	c.handler(res)
}

// New cria uma instância do consumidor
func New(url string, handler func(res *http.Response) error, cherr chan error) *Consumer {
	return &Consumer{url, handler, cherr}
}
