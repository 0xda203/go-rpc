package interfaces

// Interface Pair define os comportamentos de uma Peça
type Part interface {
	GetCode() string                       // Retorna o código da peça
	SetCode(code string)                   // Altera o código da peça
	GetName() string                       // Retorna o nome da peça
	GetDescription() string                // Retorna a descrição da peça
	GetSubcomponents() []Pair              // Retorna a lista de subcomponents da peça
	SetSubcomponents(subcomponents []Pair) // Altera a lista de subcomponentes da peça
	IsPrimitive() bool                     // Retorna se a peça é primitiva ou não (agregada)
	GetRepositoryName() string             // Retorna o nome do repositório que contém a peça
	SetRef(ref RemoteRef)                  // Retorna o nome do repositório que contém a peça
}
