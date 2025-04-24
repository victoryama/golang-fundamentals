package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	//Inicializando o Gin Gonic
	// r := gin.New() //necessita configurar tudo
	r := gin.Default() //com Logger e Recovery Middleware

	//rotas
	//ao criar a rota, configura-se o endpoint da rota
	//handler: funcao que sera executada ao chegar a requisicao
	r.GET("/victoryama/:paramPath",
		func(ctx *gin.Context) {
			ctx.String(200, "Yuji\n")
		},
		func(ctx *gin.Context) { //parametros no path (:paramPath)
			fmt.Println(ctx.Param(":paramPath"))
		},
		func(ctx *gin.Context) { //parametros no query (?paramQuery na query de requisicao)
			fmt.Println(ctx.Query("paramQuery"))
		},
		func(ctx *gin.Context) {
			ctx.String(200, "teste")
			ctx.Abort()
		},
		func(ctx *gin.Context) {
			ctx.String(200, "nao deve reproduzir")
		}) //relativePath e handlers/middlewares (quantos quiserem, executados na ordem declarada)

	err := r.Run(":9090") //anexa o router com o http.server e comeca a escutar e atender requisicoes
	//necessita especificar a porta a ser utilizado
	if err != nil {
		return
	}
}
