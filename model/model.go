// Package model contém os modelos de dados armazenados na base
package model

import "time"

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
	Cor              string `json:"nome_cor" bson:"cor"`                        // Cor do ônibus
	CriadoEm         int64  `json:"criado_em" bson:"criado_em"`                 //Data de criação do registro
	AtualizadoEm     int64  `json:"atualizado_em" bson:"atualizado_em"`         // Data de atualização do registro
	Pontos           Pontos `json:"pontos" bson:"pontos"`                       // Pontos da Linha
	Tabela           Tabela `json:"tabela" bson:"tabela"`                       // Tabela da linha
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

// Parada representa um registro da tabela de horários da linha.
type Parada struct {
	Hora     string `json:"hora" bson:"hora"`     // Hora de parada
	Ponto    string `json:"ponto" bson:"ponto"`   // Nome do ponto (UTF-8)
	Dia      string `json:"dia" bson:"dia"`       // Tipo do Dia (1 - Dia Útil, 2 - Sábado, 3 - Domingo, 4 - Feriado)
	Num      string `json:"num" bson:"num"`       // Número do ponto (de regulagem)
	Tabela   string `json:"tabela" bson:"tabela"` // Número da tabela horária
	Adaptado string `json:"adapt" bson:"adapt"`   // Informa se o ônibus tem algum tipo de adaptação para pessoas especiais
}

//Tabela representa um slice de todos os horários de parada de uma linha
type Tabela []Parada

// Veiculo representa um veículo da urbs em circulação.
type Veiculo struct {
	Codigo      string    `json:"cod" bson:"cod"`                 // Prefixo do veículo
	Refresh     string    `json:"refresh" bson:"refresh"`         // Hora da Atualização (HH24:MI)
	Latitude    string    `json:"lat" bson:"lat"`                 // Latitude (ponto flutuante)
	Longitude   string    `json:"lon" bson:"lon"`                 // Longitude (ponto flutuante)
	CodigoLinha string    `json:"codigolinha" bson:"codigolinha"` // Prefixo da Linha (3 caracteres)
	Adaptado    string    `json:"adapt" bson:"adapt"`             // Adaptado para cadeirantes (1 para sim, 0 para não)
	Tipo        string    `json:"tipo_veic" bson:"tipo_veic"`     // Tipo do ônibus {1:'COMUM',2:'SEMI PADRON', 3 :'PADRON',4 :'ARTICULADO',5 :'BIARTICULADO',6 :'MICRO', 7 :'MICRO ESPECIAL',8 :'BIARTIC. BIO',9 :'ARTIC. BIO',10:'HIBRIDO',11:'HIBRIDO BIO',12:'ELÉTRICO}'}
	Tabela      string    `json:"tabela" bson:"tabela"`           // Tabela que o veículo está executando
	Situacao    string    `json:"situacao" bson:"situacao"`       // Situação processada do veículo
	Situacao2   string    `json:"situacao2" bson:"situacao2"`     // Situação processada do veículo
	Tcount      int       `json:"tcount" bson:"tcount"`           // Contador de ciclos sem atulizar informação do veículo (informação atualizada 1)
	Sent        string    `json:"sent" bson:"sent"`               // Direção do  veículo
	CriadoEm    time.Time `json:"criado_em" bson:"criado_em"`     //data de criação do registro
	//Sentido     string `json:"sentido" bson:"sentido"`         // Sentido da linha do veículo
}

// Veiculos é um slice com os ônibus da urbs
type Veiculos []Veiculo
