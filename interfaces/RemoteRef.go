package interfaces

// Interface RemoteRef define os comportamentos de uma referência a um servidor remoto.
type RemoteRef interface {
	GetName() string    // retorna o nome do servidor remoto
	GetHost() string    // retorna o host do servidor remoto
	GetPort() string    // retorna a porta do servidor remoto
	GetAddress() string // retorna o endereço do servidor remoto no formato host:port
}
