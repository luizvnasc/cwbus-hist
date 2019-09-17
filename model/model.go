// Package model contém os modelos de dados armazenados na base
package model

//import "go.mongodb.org/mongo-driver/bson/primitive"

// Linha representa uma linha de ônibus de curitiba
//
// Foram feitas pequenas alterações na forma de apresentação do dado para
// que estes as informações da linha sejam agrupadas em um documento apenas.
//
type Linha struct {
	//ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Codigo           string `json:"cod" bson:"cod"`                             //Código da linha
	Nome             string `json:"nome" bson:"nome"`                           //  Nome da linha (UTF-8)
	SomenteCartao    string `json:"somente_cartao" bson:"somente_cartao"`       // S: Sim, N: Não, F: Finais de Semana
	CategoriaServico string `json:"categoria_servico" bson:"categoria_servico"` //Categoria da Linha (UTF-8)
	Cor              string `json:"cor" bson:"cor"`                             // Cor do ônibus
	CriadoEm         int64  `json:"criado_em" bson:"criado_em"`                 //Data de criação do registro
	AtualizadoEm     int64  `json:"atualizado_em" bson:"atualizado_em"`         // Data de atualização do registro
	Pontos           Pontos `json:"pontos" bson:"pontos"`                       // Pontos da Linha
}

// Linhas é um slice de linhas. Criado apenas para melhorar a leitura do código
type Linhas []Linha

// Ponto é um ponto de uma linha
type Ponto struct {
	Nome         string `json:"nome" bson:"nome"`                 // Nome do ponto (UTF8)
	Numero       string `json:"num" bson:"num"`                   // Número do ponto
	Latitude     string `json:"lat" bson:"lat"`                   // Latitude
	Longitude    string `json:"lon" bson:"lon"`                   // Longitude
	Sequencia    string `json:"seq" bson:"seq"`                   // Sequência do Ponto
	Grupo        string `json:"grupo" bson:"grupo"`               // Agrupadores de Pontos
	Tipo         string `json:"tipo" bson:"tipo"`                 // Tipo do Ponto (UTF8)
	Sentido      string `json:"sentido" bson:"sentido"`           // Sentido
	IDItinerario string `json:"itinerary_id" bson:"itinerary_id"` // Identificador de itinerario
}

// Pontos é um slice de Ponto
type Pontos []Ponto
