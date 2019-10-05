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

func key(name string) string {
	return prefix + name
}

func getValue(name string) string {
	return os.Getenv(key(name))
}

// ServiceURL retorna a URL dos serviços da urbs.
func ServiceURL() string {
	return getValue(serviceURL)
}

// UrbsCode retorna o código urbs de acesso aos serviços.
func UrbsCode() string {
	return getValue(urbsCode)
}

// DBStrConn retorna a string de conexão do banco de dados.
func DBStrConn() string {
	return getValue(dbStrConn)
}

// DBName retorna o nome do banco de dados.
func DBName() string {
	return getValue(dbName)
}

func WakeUpURL() string {
	return getValue(wakeUpURL)
}
