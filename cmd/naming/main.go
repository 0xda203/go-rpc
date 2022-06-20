package main

import (
	"flag"
	"go-rpc/internal/pkg/naming"
)

func main() {
	// Define as flags do executável
	// Caso elas sejam omitidas, seus valores-padrão são 127.0.0.1 (localhost)
	// e 8000, respectivamente.
	var host, port string
	flag.StringVar(&host, "host", "127.0.0.1", "host to bind to")
	flag.StringVar(&port, "port", "8000", "port to bind to")

	// Faz o parsing
	flag.Parse()

	// Inicializa serviço de nomes no host e port designados
	nameServer := new(naming.NameServer)
	nameServer.Init(host, port)
}
