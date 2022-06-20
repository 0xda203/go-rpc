package interfaces

// Interface PartRepository define os comportamentos de um repositório de peças
type PartRepository interface {
	AddPart(part Part)        // Adiciona uma Peça ao repositório de peças
	GetPart(code string) Part // Consulta uma peça pelo código no repositório e a retorna
	GetParts() []Part         // Retorna a lista de peças do repositório
}
