package main

import (
	"flag"
	"go-rpc/encoding"
	"go-rpc/internal/pkg/naming"
	s "go-rpc/internal/pkg/server"
	"go-rpc/types"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// Define as flags host, port, name e nameserver do executável
	// Caso elas sejam omitidas, seus valores-padrão são 127.0.0.1 (localhost),
	// 8001, loremipsum e 127.0.0.1:8000, respectivamente.
	var host, port, name, nameserver string

	flag.StringVar(&host, "host", "127.0.0.1", "host to bind to")
	flag.StringVar(&port, "port", "8001", "port to bind to")
	flag.StringVar(&name, "name", "loremipsum", "name to register")
	flag.StringVar(&nameserver, "ns", "127.0.0.1:8000", "nameserver address to register part repository server")

	// Faz o parsing das flags
	flag.Parse()

	// Faz o parsing do host para ver se realmente corresponde a um endereço de ip
	ip := net.ParseIP(host)

	if ip == nil {
		log.Fatalln("invalid addr", host)
	}

	// Registra tipos para correta codificação/decodificação
	encoding.RegisterConcreteTypes()

	// Inicializa repositório de peças
	partRepository := new(types.PartRepositoryImpl)

	// Inicializa servidor de repositório de peças, passando como parâmetro o reposório inicializado previamente
	partRepositoryServer := s.NewPartRepositoryServer(partRepository)

	// Inicializa servidor RPC e registra/expõe o servidor de repositório de peças como um serviço com o nome "PartRepository"
	server := rpc.NewServer()
	server.RegisterName("PartRepository", partRepositoryServer)

	// Tenta fazer a resolução do endereço host:port para checar se a porta inserida como flag já está em uso
	_, err := net.ResolveTCPAddr("tcp", host+":"+port)
	if err != nil {
		// Encerra o programa com a mensagem que a porta já está em uso
		log.Fatalf("Port already in use")
	}

	// Como o endereço do serviço de nomes é passado no formato host:port,
	// faz o split do host e da porta usando a função SPlitHostPort da biblioteca net
	// e checa por erros
	nshost, nsport, err := net.SplitHostPort(nameserver)

	if err != nil {
		log.Fatalln("Fatal error", err)
	}

	// Inicializa cliente do serviço de nomes, para resolução dos repositórios de peças
	nsclient := naming.NewNameServerClient(nshost, string(nsport))

	// Altera referência do servidor remoto do objeto partRepositoryServer
	partRepositoryServer.SetRef(types.NewRemoteRefImpl(host, port, name))

	// Começa a escutar por pacotes tcp no endereço especificado
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	// Registra o presente servidor no serviço de nomes e checa por erros
	err = nsclient.Register(host, port, name)
	if err != nil {
		log.Fatalln("Fatal error", err)
	}

	log.Printf("[!] RPC server running on %s", host+":"+port)
	log.Printf("[!] Successfully registered at nameserver with hostname %s", name)

	// Liga o servidor rpc ao socket e permite que o servidor rpc aceite
	// requisições rpc vindo desse socket.
	server.Accept(listener)
}
