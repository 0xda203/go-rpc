package types

// Estrutura RemoteRefImpl representa uma referência à um servidor.
// Implementa a interface interfaces.RemoteRef e fmt.Stringer
type RemoteRefImpl struct {
	Host string // host do servidor
	Port string // porta do servidor
	Name string // nome do servidor
}

// NewRemoteRefImpl retorna o ponteiro para uma estrutura PartImpl.
// Ela recebe como parâmetro três strings que representam o host, a porta e o nome do servidor remoto,
// respectivamente.
func NewRemoteRefImpl(host string, port string, name string) *RemoteRefImpl {
	return &RemoteRefImpl{Host: host, Port: port, Name: name}
}

// GetPort retorna o valor do atributo name da estrutura RemoteRefImpl
func (r RemoteRefImpl) GetName() string {
	return r.Name
}

// GetPort retorna o valor do atributo host da estrutura RemoteRefImpl
func (r RemoteRefImpl) GetHost() string {
	return r.Host
}

// GetPort retorna o valor do atributo port da estrutura RemoteRefImpl
func (r RemoteRefImpl) GetPort() string {
	return r.Port
}

// GetAddress retorna o endereço do servidor remoto no formato host:port
func (r RemoteRefImpl) GetAddress() string {
	return r.Host + ":" + r.Port
}

// String retorna uma string que descreve a própria estrutura RemoteRefImpl como uma string
func (r RemoteRefImpl) String() string {
	return r.GetAddress()
}
