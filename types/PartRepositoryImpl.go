package types

import "go-rpc/interfaces"

// Estrutura PartImpl representa uma repositório de peças.
// Ela implementa a interface interfaces.PartRepository.
type PartRepositoryImpl struct {
	parts []interfaces.Part // lista de peças
}

// AddPart adiciona um objeto que implementa a interface interfaces.Part à
// sua lista de peças.
func (p *PartRepositoryImpl) AddPart(part interfaces.Part) {
	p.parts = append(p.parts, part)
}

// GetPart retorna uma peça a partir do seu código.
// Ela recebe como parâmetro uma string que representa um código e faz
// uma busca linear O(n) na lista de peças, retornando-a caso seja encontrada.
// Retorna nil caso a peça não seja encontrada.
func (p *PartRepositoryImpl) GetPart(code string) interfaces.Part {
	for i := 0; i < len(p.parts); i++ {
		// Compara o código da peça ao parâmetro
		if p.parts[i].GetCode() == code {
			return p.parts[i]
		}
	}
	return nil
}

// GetParts retorna a lista de peças da estrutura PartImpl.
func (p *PartRepositoryImpl) GetParts() []interfaces.Part {
	return p.parts
}
