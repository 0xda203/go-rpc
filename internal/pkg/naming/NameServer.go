// O pacote naming fornece os tipos que definem um servidor de resolução de nomes
// genérico e básico, e um cliente para esse serviço.
package naming

import (
	"fmt"
	"go-rpc/interfaces"
	"go-rpc/types"
	"log"
	"net/http"
)

// Constantes que sinalizam respostas do serviço de nome
const (
	ERR_KEY_NOT_REGISTERED      = "key not registered"
	ERR_KEY_ALREADY_REGISTERED  = "key already registered"
	KEY_REGISTERED_SUCCESSFULLY = "success"
	ERR_POST_ONLY               = "Sorry, only POST method is supported."
	ERR_PARSING                 = "ParseForm() err"
)

// Estrutura NameServer representa um servidor http para resolução de nomes
// que utiliza uma solução simples de tabela centralizada de nome e endereço
// Name-to-address binding interna para resolução.
// As referências remotas dos servidores devem implementar a interface interfaces.RemoteRef.
type NameServer struct {
	host    string                          // host do serviço de nomes
	port    string                          // porta do serviço de nomes
	servers map[string]interfaces.RemoteRef // mapa de referências aos servidores remotos registrados.
}

// Lookup faz uma busca O(1) no mapa de referências aos servidores, retornando uma estrutura
// que implementa a interface interfaces.RemoteRef.
func (n *NameServer) lookup(key string) interfaces.RemoteRef {
	return n.servers[key]
}

// Init define os handlers para as rotas /lookup e /register do serviço de nomes
// e inicializa o servidor HTTP no host e porta designada.
func (n *NameServer) Init(host string, port string) {
	// Aloca memória para um mapa cujas chaves são strings e representam os nomes dos servidores
	// e os valores são ponteiros para objetos que implementam interfaces.RemoteRef
	n.servers = make(map[string]interfaces.RemoteRef)

	// Altera os valores das propriedades host e port para os valores recebidos por parâmetro
	n.host = host
	n.port = port

	// Define comportamento para o endpoint /lookup, que serve para fazer a resolução do endereço
	// associado a um nome
	http.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// Método POST
		case "POST":

			// Faz o parsing do formulário da requisição e sinaliza um erro caso seja encontrado
			if err := r.ParseForm(); err != nil {
				fmt.Fprint(w, ERR_PARSING)
				return
			}

			// Recupera o atributo key, que sinaliza o nome de um servidor e faz o lookup no mapa
			key := r.FormValue("key")
			ref := n.lookup(key)

			// Se a referência encontrada não for nula, escreve no corpo da resposta o endereço do servidor
			if ref != nil {
				fmt.Fprintf(w, "%s", ref.GetAddress())
			} else {
				// Caso contrário escreve que o serviço não foi encontrado
				fmt.Fprint(w, ERR_KEY_NOT_REGISTERED)
			}

		default:
			// Caso a requisição seja feita à rota por outro método sinaliza que apenas o método POST é suprotado
			fmt.Fprint(w, ERR_POST_ONLY)
		}
	})

	// Define comportamento para o endpoint /register, que serve para fazer a o registro de um servidor
	// para posterior resolução
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// Método POST
		case "POST":

			// Faz o parsing do formulário da requisição e sinaliza um erro caso seja encontrado
			err := r.ParseForm()

			if err != nil {
				fmt.Fprint(w, ERR_PARSING)
				return
			}

			// Recupera o atributo key, que sinaliza o nome de um servidor e os atributos host e port
			// que compõem o endereço do servidor que está se registrando
			key := r.FormValue("key")
			host := r.FormValue("host")
			port := r.FormValue("port")

			// Escreve false no corpo da resposta caso o nome já esteja sendo utilizada por outro servidor e retorna
			if n.lookup(key) != nil {
				fmt.Fprint(w, ERR_KEY_ALREADY_REGISTERED)
				return
			}

			// Cria uma instância de RemoteRef e define o par <chave, valor> no mapa
			n.servers[key] = types.NewRemoteRefImpl(host, port, key)

			// Faz o log do registro do servidor e escreve true no corpo da resposta
			log.Println("[!] Server at " + host + ":" + port + " registered with hostname " + key)
			fmt.Fprint(w, KEY_REGISTERED_SUCCESSFULLY)

		default:
			// Caso a requisição seja feita à rota por outro método sinaliza que apenas o método POST é suprotado
			fmt.Fprint(w, ERR_POST_ONLY)
		}
	})

	// Inicializa servidor no host e porta designadas
	log.Println("[!] HTTP server running on http://" + host + ":" + port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
