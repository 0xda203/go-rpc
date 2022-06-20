// O pacote structs fornece tipos que implementam as interfaces do pacote interfaces,
// como implementações-exemplo. Note que tipos em Go implementam interfaces implicitamente,
// bastando implementar os métodos com as suas respectivas assinaturas descritas por uma
// interface qualquer.
package types

import (
	"fmt"
	"go-rpc/interfaces"
)

// Estrutura PairImpl representa um par (elemento, quantidade).
// Ela implementa as interfaces interfaces.Pair e fmt.Stringer (que define o método String() que altera o comportamento do print, útil para o propósito de log e debug, equivalente ao toString()).
type PairImpl struct {
	SubPart  interfaces.Part // peça do par
	Quantity int             // quantidada peça
}

// NewPairImpl retorna o ponteiro para uma estrutura PairImpl.
// Ela recebe como parâmetro um ponteiro para uma Peça de qualquer tipo e um valor do tipo inteiro.
func NewPairImpl(subPart interfaces.Part, quantity int) *PairImpl {
	return &PairImpl{SubPart: subPart, Quantity: quantity}
}

// GetElement retorna a propriedade SubPart da estrutura PairImpl.
func (p PairImpl) GetPart() interfaces.Part {
	return p.SubPart
}

// GetQuantity retorna a propriedade Quantity da estrutura PairImpl.
func (p PairImpl) GetQuantity() int {
	return p.Quantity
}

// SetQuantity altera o valor da propriedade Quantity da estrutura PairImpl.
// Ela aceita como parâmetro um valor do tipo inteiro.
func (p *PairImpl) SetQuantity(quantity int) {
	p.Quantity = quantity
}

// String retorna uma string que descreve a própria estrutura PairImpl como uma string
func (p PairImpl) String() string {
	// %#v mostra a estrutura com os atributos e seus respectivos valores
	return fmt.Sprintf("Par{%s, Quantidade: %d]}", p.SubPart, p.GetQuantity())
}
