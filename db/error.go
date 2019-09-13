package db

// Error representa um erro que pode ocorrer durante a tentativa de se coectar com um banco de dados.
type Error string

// Error implementa a interface Error
func (e Error) Error() string {
	return string(e)
}

const (
	//ErrNoConnString : String de conexão não informado
	ErrNoConnString = Error("String de conexão não informado")
)
