// Package model contém os modelos de dados armazenados na base
package model

//import "go.mongodb.org/mongo-driver/bson/primitive"

// Linha representa uma linha de ônibus de curitiba
//
// Foram feitas pequenas alterações na forma de apresentação do dado para
// que estes tenham tipos definidos e não tratados apenas como strings que
// é como o serviço /getLinhas retorna as informações.
//
// Exemplo de retorno do serviço:
//
//
// [
// 	{
// 		COD: "464",
// 		NOME: "A. MUNHOZ / J. BOTÂNICO",
// 		SOMENTE_CARTAO: "S",
// 		CATEGORIA_SERVICO: "CONVENCIONAL",
// 		NOME_COR: "AMARELA"
// 	},
// ...
// ]
type Linha struct {
	//ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Codigo           int    `json:"cod" bson:"cod"`
	Nome             string `json:"nome" bson:"nome"`
	SomenteCartao    string   `json:"somente_cartao" bson:"somente_cartao"`
	CategoriaServico string `json:"categoria_servico" bson:"categoria_servico"`
	Cor              string `json:"cor" bson:"cor"`
	CriadoEm 		 int64  `json:"criado_em" bson:"criado_em"`
	AtualizadoEm 	 int64  `json:"atualizado_em" bson:"atualizado_em"`
}

// Linhas é um vetor de linhas. Criado apenas para melhorar a leitura do código
type Linhas []Linha
