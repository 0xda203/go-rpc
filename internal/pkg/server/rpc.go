// O pacote server fornece o tipo que implementa a definição de um servidor que implementa
// um objeto repositório de peças que será exposto via RPC, como definido na descrição do exercício-programa.
package server

import (
	"go-rpc/interfaces"

	"github.com/google/uuid"
)

// Estrutura PartRepositoryServer representa um servidor que implementa um objeto PartRepository.
// Essa é a estrutura do objeto que será registrada e exposta via RPC.

// Por definição, apenas métodos que satisfazem o seguinte critério podem ser expostos para acesso
// remoto utilizando a biblioteca rpc da linguagem Go:
//
// - O tipo do método (ou da a estutura que a implementa) deve ser exportado (i.e. seu nome é capitalizado e seus campos são públicos).
// - O método deve ser público. Nota-se que em Go um método é público quando seu nome é capitalizado.
// - O método tem dois argumentos, ambos exportados (ou públicos).
// - O segundo argumento do método é um ponteiro.
// - O tipo do valor de retorno precisa ser error.
//
// Esquematicamente, a assinatura do método deve parecer semanticamente com:
//
// func (t *T) MethodName(argType T1, replyType *T2) error
//
// onde T1 e T2 podem ser serializados através do gob, o formato do Go de serializar e desserializar dados em formato binário.
// Gob é encontrada em encoding/gob e é similar ao "Picke" do Python ou "Serialization" do Java.
// Diferente de json ou xml, porém, é um formato específico para comunicação entre um servidor e cliente escritos em Go,
// o que torna a comunicação muito mais eficiente e otimizada.

// Esses requisitos se aplicam mesmo se um codec diferente for usado.

// O primeiro argumento do método representa os argumentos fornecidos pelo chamador;
// o segundo argumento representa os parâmetros de resultado a serem retornados ao chamador.
// O valor de retorno do método, se não for nulo, é passado de volta como uma string que o cliente
// vê como se tivesse sido criada por errors.New. Se um erro for retornado, o parâmetro de resposta
// não será enviado de volta ao cliente.
type PartRepositoryServer struct {
	partRepository interfaces.PartRepository // objeto PartRepository
	ref            interfaces.RemoteRef      // referência do servidor remoto
}

// NewPartRepositoryServer retorna o ponteiro para uma estrutura PartRepositoryServer.
// Ela recebe como parâmetro um objeto que implementa a interface interfaces.PartRepository.
func NewPartRepositoryServer(p interfaces.PartRepository) *PartRepositoryServer {
	return &PartRepositoryServer{partRepository: p}
}

// AddPart adiciona uma peça ao repositório de peças da estrutura PartRepositoryServer.
// Ela atribui um identificador único à peça e define a sua referência ao servidor remoto como a
// referência remota do próprio servidor.
// Recebe como parâmetros um ponteiro para uma peça, que será adicionada à lista, e um ponteiro para uma peça, que passará
// a apontar à própria peça inserida, após definir o valor identificador e a referência remota do servidor.
// Retorna por padrão nulo, sinalizando que não houve erro na comunicação.
func (p *PartRepositoryServer) AddPart(part *interfaces.Part, reply *interfaces.Part) error {
	// Gera novo identificador
	id := uuid.New().String()

	// Altera o código do objeto e a referência ao servidor
	(*part).SetCode(id)
	(*part).SetRef(p.ref)

	// Adiciona a peça usando a API do objeto PartRepository
	p.partRepository.AddPart(*part)

	// Armazena no segundo parâmetro o endereço de memória peça adicionada
	*reply = *part

	return nil
}

// GetPart consulta uma peça a partir de um código no repositório de peças e a retorna ao usuário.
// Recebe como parâmetros uma string, que indica o código da peça a ser buscada
// e um ponteiro para uma peça, que passará a apontar para a peça buscada, caso seja encontrada.
// Retorna por padrão nulo, sinalizando que não houve erro na comunicação.
func (p *PartRepositoryServer) GetPart(code string, out *interfaces.Part) error {
	*out = p.partRepository.GetPart(code)
	return nil
}

// GetParts retorna a lista de peças do repositório de peças.
// Recebe como parâmetros um ponteiro para uma string dummy, que será ignorado (é dessa forma para atender aos critérios previamente mencionados),
// e um ponteiro para uma lista de peça, na qual será armazenada, que passará a apontar para a própria lista de peças do objeto PartRepository.
// Retorna por padrão nulo, sinalizando que não houve erro na comunicação.
func (p *PartRepositoryServer) GetParts(_ string, out *[]interfaces.Part) error {
	*out = p.partRepository.GetParts()
	return nil
}

// SetRef altera o valor da propriedade ref da estrutura PartRepositoryServer.
// Ela recebe como parâmetro uma referência a um servidor remoto que implementa a interface interfaces.RemoteRef
// Esse método não atende aos critérios previamente citados e, portanto, não é registrado e exposta via RPC.
func (p *PartRepositoryServer) SetRef(ref interfaces.RemoteRef) {
	p.ref = ref
}
