package naming

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Estrutura NameServerCliente representa um cliente para o servidor http de resolução de nomes
// NameServer.
type NameServerClient struct {
	host string // host do serviço de nomes
	port string // porta do serviço de nomes
}

// NewNameServerClient retorna o ponteiro para uma estrutura NameServerClient.
// Ela recebe como parâmetro duas strings que representam o host e a porta do serviço de nomes remoto, respectivamente.
func NewNameServerClient(host string, port string) *NameServerClient {
	return &NameServerClient{host, port}
}

// Lookup faz uma consulta ao serviço de nomes a fim de resolver o endereço a partir do nome do servidor
// Recebe como parâmetro uma string chave, que sinaliza o nome do serviço a ser resolvido, e retorna
// uma string e um erro, que sinalizam, respectivamente, o endereço do servidor, caso seja resolvido,
// e um erro, caso o nome não tenha sido resolvido, sinalizando que a chave provavelmente não foi registrada.
func (n *NameServerClient) Lookup(key string) (string, error) {
	// Prepara corpo da requisição
	data := url.Values{
		"key": {key},
	}

	// Faz a requisição ao serviço de nomes conectado no endpoint /lookup e faz o logging caso haja um erro na requisição
	resp, err := http.PostForm(n.getAddress()+"/lookup", data)
	if err != nil {
		log.Fatal(err)
	}

	// Faz a leitura do corpo da requisição e sinaliza caso haja algum erro
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// Converte o corpo da string para o tipo string e checa por erros
	sb := string(body)

	switch sb {
	case ERR_KEY_NOT_REGISTERED, ERR_PARSING:
		return "", errors.New(sb)
	default:
		// retorna o endereço normalmente
		return sb, nil
	}
}

// Register faz uma requisição ao serviço de nomes a fim de registrar o endereço de origem à um nome.
// Recebe como parâmetro o host, porta e nome do servidor que quer se registar no serviço de nomes, respectivamente, e retorna
// um  um erro, que sinaliza uma falha no registro, caso ocorra.
func (n *NameServerClient) Register(host string, port string, name string) error {
	// Prepara corpo da requisição contendo dados para registro
	data := url.Values{
		"host": {host},
		"port": {port},
		"key":  {name},
	}

	// Faz a requisição ao servidor no endpoint /register e sinaliza caso ocorra algum erro
	resp, err := http.PostForm(n.getAddress()+"/register", data)

	if err != nil {
		log.Fatal(err)
	}

	// Faz a leitura do corpo da requisição e sinaliza caso ocorra algum erro
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// Converte o corpo da string para o tipo string e checa por erros
	sb := string(body)

	switch sb {
	case KEY_REGISTERED_SUCCESSFULLY:
		return nil
	default:
		return errors.New(sb)
	}
}

// getAddress retorna o endereço no formato host:porta do serviço de nomes
func (n *NameServerClient) getAddress() string {
	return "http://" + n.host + ":" + n.port
}
