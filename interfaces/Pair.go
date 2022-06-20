// O pacote interfaces fornece interfaces básicas para as estruturas
// utilizadas no projeto.
package interfaces

// Interface Pair define os comportamentos de um Par (peça, quantidade)
type Pair interface {
	GetPart() Part            // Retorna a peça do par
	GetQuantity() int         // Retorna a quantidade do par
	SetQuantity(quantity int) // Altera a propriedade quantidade do par
}
