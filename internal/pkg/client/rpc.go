// O pacote client fornece o tipo que implementa a definição de um cliente do servidor que implementa
// um objeto repositório de peças, como definido na descrição do exercício-programa.
package client

import (
	"go-rpc/interfaces"
	"log"
	"net/rpc"
)

// Estrutura PartRepositoryClient representa um cliente do servidor servidor PartRepositoryServer.
// Essa é a estrutura do objeto que será registrada e exposta via RPC. Em geral, ela converte
// uma chamada local p.AddPart numa chamada rpc definida pela API net/rpc.
type PartRepositoryClient struct {
	ref    interfaces.RemoteRef
	client *rpc.Client // ponteiro para cliente RPC
}

// NewPartRepositoryClient retorna o ponteiro para uma estrutura PartRepositoryClient.
// Ela recebe como parâmetro um ponteiro para um cliente RPC.
func NewPartRepositoryClient(client *rpc.Client, ref interfaces.RemoteRef) *PartRepositoryClient {
	return &PartRepositoryClient{ref, client}
}

// AddPart adiciona uma peça ao repositório de peças.
// Ela recebe como parâmetro um objeto que implementa a interface interfaces.Part
// e faz uma chamada RPC via cliente RPC passando o ponteiro para a peça a ser
// adicionada e o ponteiro para a peça a ser devolvida com informações adicionais, respectivamente.
// Ela retorna o objeto devolvido, que implementa a interface interfaces.Part.
func (p *PartRepositoryClient) AddPart(part interfaces.Part) interfaces.Part {
	// Faz chamada RPC
	err := p.client.Call("PartRepository.AddPart", &part, &part)
	// Sinaliza erros no canal
	if err != nil {
		log.Fatal("Fatal error:", err)
	}
	return part
}

// GetPart recupera uma peça do repositório de peças a partir do seu código.
// Ela recebe como parâmetro uma string contendo o código a ser buscado
// e faz uma chamada RPC via cliente RPC passando o ponteiro para o código a ser buscado
// o ponteiro para a peça a ser devolvida, caso encontrada, respectivamente.
// Ela retorna o objeto devolvido, que implementa a interface interfaces.Part.
func (p *PartRepositoryClient) GetPart(code string) interfaces.Part {
	var part interfaces.Part
	// Faz chamada RPC
	err := p.client.Call("PartRepository.GetPart", &code, &part)
	// SInaliza erros no canal
	if err != nil {
		log.Fatal("Fatal error:", err)
	}
	return part
}

// GetParts recupera a lista de peças do repositório de peças a partir do seu código.
// Ela faz uma chamada RPC via cliente RPC passando o ponteiro uma string dummy, que será ignorada,
// e um ponteiro para a lista de peças a ser devolvida, respectivamente.
// Ela retorna uma lista de objetos que implementam a interface interfaces.Part.
func (p *PartRepositoryClient) GetParts() []interfaces.Part {
	var parts []interfaces.Part
	// Faz chamada RPC
	err := p.client.Call("PartRepository.GetParts", "dummy", &parts)
	// Sinaliza erros no canal
	if err != nil {
		log.Fatal("RPC error:", err)
	}
	return parts
}

// GetRepositoryName retorna o nome do repositório de peças atualmente conectado
func (p PartRepositoryClient) GetRepositoryName() string {
	return p.ref.GetName()
}
