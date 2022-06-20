package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-rpc/encoding"
	"go-rpc/interfaces"
	"go-rpc/internal/pkg/client"
	"go-rpc/internal/pkg/naming"
	"go-rpc/types"
	"log"
	"net"
	"net/rpc"
	"os"
	"strconv"
)

var isConnected bool = false

var nsClient *naming.NameServerClient
var rpcClient *rpc.Client
var currentRepo *client.PartRepositoryClient // repositório corrente
var currentPart interfaces.Part              // peça corrente
var currentSubcomponents []interfaces.Pair   // lista de sub-peças corrente

// addsupart adiciona à lista de sub-peças n unidades da peça corrente.
// Recebe como parâmetro um objeto que implementa a interface interfaces.Part e
// um inteiro que representa a quantidade de unidades desse objeto.
func addsubpart(subPart interfaces.Part, n int) {
	// Checa se subPart já está na lista
	for i := 0; i < len(currentSubcomponents); i++ {
		curr := currentSubcomponents[i].GetPart()
		if (curr).GetCode() == (subPart).GetCode() {
			// Apenas altera a quantidade do componente na lista
			currentSubcomponents[i].SetQuantity(currentSubcomponents[i].GetQuantity() +
				n)
			return
		}
	}

	// Adiciona o novo par na lista de subcomponentes corrent
	currentSubcomponents = append(currentSubcomponents, types.NewPairImpl(subPart, n))
}

// listp lista as peças do repositório corrente
func listp() {
	parts := currentRepo.GetParts()

	if len(parts) == 0 {
		fmt.Printf("[!] Lista de peças vazia")
	} else {
		fmt.Printf("[!] Peças do repositório corrente {%s}:\n", currentRepo.GetRepositoryName())
	}

	for i := 0; i < len(parts); i++ {
		fmt.Printf("\t%v", parts[i])
		if i < len(parts)-1 {
			fmt.Println()
		}
	}

}

// bind tenta resolver através do cliente de serviço de nomes o servidor de repositório
// de peças a partir de um nome e retorna o ponteiro para uma estrutura client.PartRepositoryClient
// que fornece a API para interagir com o servidor remoto.
// Recebe como parâmetro o nome do servidor para se conectar.
func bind(serverName string) *client.PartRepositoryClient {
	// Caso o cliente com o nameserver seja nulo, retorne nulo
	if nsClient == nil {
		return nil
	}

	// Tenta fazer o lookup através do cliente do serviço de nomes
	// usando o nome de servidor passado como parâmetro
	address, err := nsClient.Lookup(serverName)

	// Faz o logging de erro se for o caso
	if err != nil {
		fmt.Printf("[!] Oops, servidor com nome %s não encontrado\n", serverName)
		return nil
	}

	// Tenta se conectar ao endereço do servidor rpc e reporta erros
	// caso ocorram
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Connection error:", err)
	}

	isConnected = true

	// inicializa um novo cliente rpc
	rpcClient = rpc.NewClient(conn)

	// Separa o endereço em host e port para inicializar um objeto RemoteRef
	host, port, _ := net.SplitHostPort(address)

	fmt.Printf("[!] Conectado com sucesso ao servidor %s{%s}", serverName, address)
	// Retorna nova estrutura client.PartRepositoryClient construída a partir do cliente RPC
	return client.NewPartRepositoryClient(rpcClient,
		types.NewRemoteRefImpl(host, port, serverName))
}

// quit encerra a execução do cliente e finaliza a aplicação
func quit() {
	fmt.Printf("[!] Cliente finalizado.\n")
	// finaliza a aplicação
	os.Exit(0)
}

func readline(scanner *bufio.Scanner) string {
	var strInput string

	if scanner.Scan() {
		strInput = scanner.Text()
	}

	return strInput
}

func main() {

	// Define a flag nameserver do executável
	// Caso ela seja omitida, seu valor-padrão é 127.0.0.1:8000.
	var nameserver string
	flag.StringVar(&nameserver, "ns", "127.0.0.1:8000", "nameserver address to key resolution")

	// Faz o parsing das flags
	flag.Parse()

	// Como o endereço do serviço de nomes é passado no formato host:port,
	// faz o split do host e da porta usando a função SPlitHostPort da biblioteca net
	// e checa por erros
	host, port, err := net.SplitHostPort(nameserver)

	if err != nil {
		log.Fatalln("Fatal error", err)
	}

	// Faz o parsing do host para ver se realmente corresponde a um endereço de ip
	ip := net.ParseIP(host)

	// Encerra o programa com a mensagem de erro invalid addr caso não seja um ip válido
	if ip == nil {
		log.Fatalln("invalid addr", host)
	}

	// Registra tipos para correta codificação/decodificação
	encoding.RegisterConcreteTypes()

	// Inicializa o cliente do serviço de nomes
	nsClient = naming.NewNameServerClient(host, port)

	var repoName string
	for !isConnected {
		fmt.Printf("[!] Digite o nome do repositório para se conectar: ")
		fmt.Scanln(&repoName)
		currentRepo = bind(repoName)
	}

	var command string
	fmt.Printf("\n[!] Comandos disponíveis: (bind|listp|getp|showp|clearlist|addsubpart|addp|quit)")

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Printf("\n > ")
		command = readline(scanner)
		switch command {
		case "bind":
			fmt.Printf("[!] Digite o nome do repositório para se conectar: ")
			repoName = readline(scanner)
			currentRepo = bind(repoName)
			for currentRepo == nil {
				fmt.Printf("[!] Digite o nome do repositório para se conectar: ")
				repoName = readline(scanner)
				currentRepo = bind(repoName)
			}
		case "listp":
			listp()
		case "getp":
			fmt.Printf("[!] Digite o código da peça para busca: ")
			partCode := readline(scanner)
			part := currentRepo.GetPart(partCode)
			if part == nil {
				fmt.Printf("[!] Peça com código %s não encontrada.", partCode)
			} else {
				currentPart = part
				fmt.Printf("[!] Peça corrente definida como %v.", currentPart)
			}
		case "showp":
			if currentPart == nil {
				fmt.Printf("[!] Peça corrente não foi definida.")
			} else {
				fmt.Printf("[!] Peça corrente: %s", currentPart)
			}
		case "clearlist":
			// Note que em Go, atribuir o valor nil a um slice é equivalente a esvaziá-lo.
			currentSubcomponents = nil
			fmt.Printf("[!] Lista de sub-peças corrente limpa com sucesso.")
		case "addsubpart":
			if currentPart == nil {
				fmt.Printf("[!] Peça corrente ainda não foi definida.")
			} else {
				fmt.Printf("[!] Digite a quantidade n da peça corrente que deseja adicionar: ")
				n := readline(scanner)
				no, err := strconv.Atoi(n)
				if err == nil {
					addsubpart(currentPart, int(no))
					fmt.Printf("[!] Peça adicionada à lista de sub-peças com sucesso.")
				} else {
					fmt.Printf("[!] Por favor, digite um número.")
				}
			}
		case "addp":
			fmt.Printf("[!] Digite o nome da peça: ")
			name := readline(scanner)
			fmt.Printf("[!] Digite a descrição da peça: ")
			description := readline(scanner)
			newPart := types.NewPartImpl(name, description)
			newPart.SetSubcomponents(currentSubcomponents)
			p := currentRepo.AddPart(newPart)
			fmt.Printf("[!] Peça adicionada com sucesso (código=%s)", p.GetCode())
		case "quit":
			quit()
		default:
			fmt.Printf("[!] Comando não reconhecido.")
		}
	}

}
