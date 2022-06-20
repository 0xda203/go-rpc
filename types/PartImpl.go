package types

import (
	"fmt"
	"go-rpc/interfaces"
)

// Estrutura PartImpl representa uma Peça.
// Ela implementa as interfaces interfaces.Part e fmt.Stringer.
type PartImpl struct {
	Code          string               // código da peça
	Name          string               // nome da peça
	Description   string               // descrição da peça
	Subcomponents []interfaces.Pair    // lista de subcomponents (i.e. qualquer estrutura que implementa interfaces.Pair) da peça
	Ref           interfaces.RemoteRef // referência do servidor que contém a peça
}

// NewPartImpl retorna o ponteiro para uma estrutura PartImpl.
// Ela recebe como parâmetro duas strings que representam o nome e a descrição da peça,
// respectivamente. O atributo subcomponents é inicializado como uma lista vazia por definição.
func NewPartImpl(name string, description string) *PartImpl {
	return &PartImpl{Name: name, Description: description}
}

// GetCode retorna a propriedade code da estrutura PartImpl.
func (p PartImpl) GetCode() string {
	return p.Code
}

// GetName retorna a propriedade name da estrutura PartImpl.
func (p PartImpl) GetName() string {
	return p.Name
}

// GetDescription retorna a propriedade description da estrutura PartImpl.
func (p PartImpl) GetDescription() string {
	return p.Description
}

// GetSubcomponents retorna a propriedade subcomponents da estrutura PartImpl.
func (p PartImpl) GetSubcomponents() []interfaces.Pair {
	return p.Subcomponents
}

// SetSubcomponents altera o valor da propriedade subcomponents da estrutura PartImpl.
// Ela aceita como parâmetro um slice de elementos que implementam a interface interfaces.Pair.
// Caso o parâmetro subcomponents seja um valor nulo, o slice é limpado.
func (p *PartImpl) SetSubcomponents(subcomponents []interfaces.Pair) {
	// Em Go, alterar o valor de um slice para nil é equivalente a limpá-lo.
	p.Subcomponents = subcomponents
}

// IsPrimitive retorna true se a peça é primitiva (i.e. sem subcomponentes)
// ou false caso a peça seja agregada.
func (p PartImpl) IsPrimitive() bool {
	return len(p.Subcomponents) == 0
}

// GetRepositoryName retorna o nome do servidor que contém a peça
func (p PartImpl) GetRepositoryName() string {
	return p.Ref.GetName()
}

// SetCode altera o valor da propriedade code da estrutura PartImpl.
// Ela aceita como parâmetro uma string.
func (p *PartImpl) SetCode(code string) {
	p.Code = code
}

// SetCode altera o valor da propriedade ref da estrutura PartImpl.
// Ela aceita como parâmetro um objeto que implementa a interface interfaces.RemoteRef.
func (p *PartImpl) SetRef(ref interfaces.RemoteRef) {
	p.Ref = ref
}

// String retorna uma string que descreve a própria estrutura PairImpl como uma string
func (p PartImpl) String() string {
	// %#v mostra a estrutura com os atributos e seus respectivos valores
	ptype := "Primitivo"
	if len(p.GetSubcomponents()) > 0 {
		ptype = "Agregada"
	}
	str := fmt.Sprintf(`Peça{Código:%s, Repositório: %s, Tipo: %s, Nome:%s, Descrição:%s, Subpeças (%d): [`, p.Code, p.GetRepositoryName(), ptype, p.Name, p.Description, len(p.GetSubcomponents()))
	for i, subPart := range p.GetSubcomponents() {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%v", subPart)
	}
	return str + "]}"
}
