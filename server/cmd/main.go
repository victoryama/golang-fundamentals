package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// Server 2
// conecta uma funcao handler a URLs de entrada cujo path comeca com /
// conecta com qualquer URL e inicia um servidor que fica ouvindo requisicoes de entrada na porta 8000
// requisicao tem a estrutura http.requests com campos como
// URL da requisicao de entrada
func main() {
	http.HandleFunc("/", handler) //cada requisicao se chama handler
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ecoa o componente Path do URL requisitado
// quando uma requisicao chega eh enviada para a funcao handler
// funcao handler extrai o componente path do URL de requisicao
// e envia como resposa usando o fmt.Fprintf
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock() //permite que apenas uma gorrotine acesse a variavel em determinado instante
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//servidor tem 2 handlers, URL requisitado determina qual deles [e chamado
// /count para counter e os demais para handler

// counter ecoa o numero de chamadas ate agora
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Println(2, "Count %d\n", count)
	mu.Unlock()
}
