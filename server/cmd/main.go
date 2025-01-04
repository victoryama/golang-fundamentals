package main

import (
	"fmt"
	"log"
	"net/http"
)

// Server 1
// conecta uma funcao handler a URLs de entrada cujo path comeca com /
// conecta com qualquer URL e inicia um servidor que fica ouvindo requisicoes de entrada na porta 8000
// requisicao tem a estrutura http.requests com campos como
// URL da requisicao de entrada
func main() {
	http.HandleFunc("/", handler) //cada requisicao se chama handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ecoa o componente Path do URL requisitado
// quando uma requisicao chega eh enviada para a funcao handler
// funcao handler extrai o componente path do URL de requisicao
// e envia como resposa usando o fmt.Fprintf
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
