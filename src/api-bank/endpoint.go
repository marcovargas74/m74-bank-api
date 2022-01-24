package m74bankapi

import (
	"fmt"
	"log"
	"net/http"
)

const (
	serverPort = ":5000"
)

func getPlayerPoints(name string) string {

	if name == "Maria" {
		return "20"
	}

	if name == "Pedro" {
		return "10"
	}
	return ""

}

//TIPO TRATADOR CRIASSE A INTERFACE A STRUCT E O INICIALISADOR DO SERVER
type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
	RegistrarVitoria(nome string)
}

type ServidorJogador struct {
	Armazenamento ArmazenamentoJogador
	//Roteador      *http.ServeMux
	http.Handler
}

//NovoServidorJogador Cria Servidor
func NovoServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {

	s := new(ServidorJogador)
	s.Armazenamento = armazenamento

	roteador := http.NewServeMux()
	roteador.Handle("/liga", http.HandlerFunc(s.manipulaLiga))
	roteador.Handle("/jogadores/", http.HandlerFunc(s.manipulaJogadores))

	s.Handler = roteador

	return s

}

func (s *ServidorJogador) mostrarPontuacao(w http.ResponseWriter, jogador string) {

	pontuacao := s.Armazenamento.ObterPontuacaoJogador(jogador)

	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}

func (s *ServidorJogador) registrarVitoria(w http.ResponseWriter, jogador string) {
	s.Armazenamento.RegistrarVitoria(jogador)
	w.WriteHeader(http.StatusAccepted)
}

func (s *ServidorJogador) manipulaLiga(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *ServidorJogador) manipulaJogadores(w http.ResponseWriter, r *http.Request) {

	jogador := r.URL.Path[len("/jogadores/"):]

	switch r.Method {
	case http.MethodPost:
		s.registrarVitoria(w, jogador)
	case http.MethodGet:
		s.mostrarPontuacao(w, jogador)
	}
}

/*
func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	s.Roteador.ServeHTTP(w, r)
	/*roteador := http.NewServeMux()

	roteador.Handle("/liga", http.HandlerFunc(s.manipulaLiga))
	roteador.Handle("/jogadores/", http.HandlerFunc(s.manipulaJogadores))*/

/*fmt.Printf("URL in %v  MEthod%v\n", r.URL, r.Method)
	roteador.Handle("/liga", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	roteador.Handle("/jogadores/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		jogador := r.URL.Path[len("/jogadores/"):]

		switch r.Method {
		case http.MethodPost:
			s.registrarVitoria(w, jogador)
		case http.MethodGet:
			s.mostrarPontuacao(w, jogador)
		}
	}))* /

	//roteador.ServeHTTP(w, r)
}
*/

func DefaultEndpoint(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Default data in %v\n", r.URL)
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	fmt.Fprint(w, "Endpoint not found")
}

/*
 * BANK INICIA AQUI
 */
func callbackAccount(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("callbackAccount data in %v\n", r.URL)
	fmt.Fprint(w, message)
}

func callbackLogin(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("callbackLogin data in %v\n", r.URL)
	fmt.Fprint(w, message)
}

func callbackTransfer(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("callbackTransfer data in %v\n", r.URL)
	fmt.Fprint(w, message)
}

//StartAPI inicia o servidor http
func StartAPI(modo string) {
	//tratador := http.HandlerFunc(ServidorJogador)
	//log.Fatal(http.ListenAndServe(serverPort, tratador))
	HandleFuncions()
	log.Fatal(http.ListenAndServe(serverPort, nil))
}

//HandleFuncions Inclui os endpoint
func HandleFuncions() {
	http.HandleFunc("/", DefaultEndpoint)
	//http.HandleFunc("/jogadores/Maria", ServidorJogador)
	//http.HandleFunc("/jogadores/Pedro", ServidorJogador)

	//*TODO endpoint usado no banc
	http.HandleFunc("/accounts", callbackAccount)
	http.HandleFunc("/login", callbackLogin)
	http.HandleFunc("/transfers", callbackTransfer)
}

// INLCUIR EM UM OUTRO ARQUIVO DE PERSISTENCIA

// NovoArmazenamentoJogadorEmMemoria cria um ArmazenamentoJogador vazio
func NovoArmazenamentoJogadorEmMemoria() *ArmazenamentoJogadorEmMemoria {
	return &ArmazenamentoJogadorEmMemoria{map[string]int{}}
}

// ArmazenamentoJogadorEmMemoria armazena na memória os dados sobre os jogadores
type ArmazenamentoJogadorEmMemoria struct {
	armazenamento map[string]int
}

// RegistrarVitoria irá registrar uma vitoria
func (a *ArmazenamentoJogadorEmMemoria) RegistrarVitoria(nome string) {
	a.armazenamento[nome]++
}

// ObterPontuacaoJogador obtém as pontuações para um jogador
func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return a.armazenamento[nome]
}
