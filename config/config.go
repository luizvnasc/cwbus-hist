// Package config obtém as configurações da app que estão nas variáveis de ambiente
package config

import "os"

const (
	prefix     = "CWBUS_"
	serviceURL = "URBS_SERVICE_URL"
	urbsCode   = "URBS_CODE"
	dbStrConn  = "DB_URL"
	dbName     = "DB_HIST"
	wakeUpURL  = "WAKEUP_URL"
)

// Configurer é a interface que define um configurador no sistema.
type Configurer interface {
	ServiceURL() string
	UrbsCode() string
	DBName() string
	DBStrConn() string
	WakeUpURL() string
}

// EnvConfigurer é um confiurador que  capitura as configurações das variáveis de ambiente.
type EnvConfigurer struct{}

func (ec EnvConfigurer) key(name string) string {
	return prefix + name
}

func (ec EnvConfigurer) getValue(name string) string {
	return os.Getenv(ec.key(name))
}

// ServiceURL retorna a URL dos serviços da urbs.
func (ec EnvConfigurer) ServiceURL() string {
	return ec.getValue(serviceURL)
}

// UrbsCode retorna o código urbs de acesso aos serviços.
func (ec EnvConfigurer) UrbsCode() string {
	return ec.getValue(urbsCode)
}

// DBStrConn retorna a string de conexão do banco de dados.
func (ec EnvConfigurer) DBStrConn() string {
	return ec.getValue(dbStrConn)
}

// DBName retorna o nome do banco de dados.
func (ec EnvConfigurer) DBName() string {
	return ec.getValue(dbName)
}

// WakeUpURL retorna a url utilizada para acordar o dyno do heroku
func (ec EnvConfigurer) WakeUpURL() string {
	return ec.getValue(wakeUpURL)
}
