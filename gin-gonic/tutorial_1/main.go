package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type user struct {
	Name string `json:"name" validate:"required"`              //especifica o local do valor no json
	Age  int32  `json:"age" validate:"required,min=0,max=130"` //colocando o nome do campo do json
	//especifica a validacao
}

func main() {

	// gRouter := gin.New() //inicializando o gin cru, necessita configurar
	gRouter := gin.Default() //apresenta mais recursos automaticamente
	// recebe o router do GIN

	gRouter.GET("/ping", func(c *gin.Context) { //trata a requisicao GET que seja "/ping"
		//pega o c q é o contexto do gin
		c.JSON(200, gin.H{ //retorna o codigo 200 e a mensagem em message pong em um Json
			"message": "pong",
		})
	})

	//path param
	gRouter.GET("/user/:name/", func(c *gin.Context) { //path param a partir do :name, URL tem o valor
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//querystring
	gRouter.GET("/welcome", func(c *gin.Context) { //passar nome das variaveis e o valor
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	//body params / binding from json
	gRouter.POST("/welcome", func(c *gin.Context) {
		var user user
		err := c.ShouldBindBodyWithJSON(&user) //passa ponteiro para user, para pegar as variaveis e passar para o ponteiro user
		if err != nil {
			c.JSON(http.StatusBadRequest, "Erro ao converter dados")
			return
		}
		validate := validator.New()
		err = validate.Struct(user) //solicita a validacao do user
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Erro ao validar dados")
			return
		}

		c.JSON(http.StatusOK, user) //gin converte o user para JSON
	})

	gRouter.Run(":8081")

}
