package encoding

import (
	"encoding/gob"
	"go-rpc/types"
)

// RegisterConcreteTypes registra os tipos concretos das interfaces
// utilizadas na aplicação para correta codificação/decodificação
// entre gobs e os tipos concretos.
func RegisterConcreteTypes() {
	gob.Register(types.RemoteRefImpl{})
	gob.Register(&types.PartImpl{})
	gob.Register(&types.PairImpl{})

}
