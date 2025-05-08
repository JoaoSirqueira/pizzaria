package main

import (
	"pizzaria/internal/data"

	"github.com/gin-gonic/gin"
	"pizzaria/internal/handler"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()
	// Criando uma rota
	router.GET("/pizzas", handler.GetPizzas)
	router.POST("/pizzas", handler.PostPizzas)
	router.GET("/pizzas/:id", handler.GetPizzasByID)
	router.DELETE("/pizzas/:id", handler.DeletePizzaById) // Deletar uma pizza
	router.PUT("/pizzas/:id", handler.UpdatePizzaByID)    // Editar ou atualizar uma pizza
	router.POST("pizzas/:id/reviews", handler.PostReview)
	router.Run()                                   // Executar rota
}

// underline ali no FOR significa que vai ocultar o índice

// executar a operação: go run . OU go run main.go
// trazer dependências e pacotes externos: go get

// dependência do gin:
// go get github.com/gin-gonic/gin

// utilizar atributos de uma struct com letra minúscula significa que eles são privados, ou seja, não pode acessar essas informações de outros lugares
// utilizar atributos de uma struct com a primeira letra maiúscula significa que eles são públicos, ou seja, pode ser acessados essas informações de outros lugares.
// Quando trabalhamos com API é comum deixar as coisas com minúsculos para fazer com o GO usa-se um highlight e coloca o código: `json:"id"` exemplo:
// type Pizza struct {
//     ID    int     `json:"id"`
//     Nome  string  `json:"nome"`
//     Preco float64 `json:"preco"`
// }

// A função Decode é utilizada para decodificar dados JSON e carregá-los em uma estrutura Go
// enquanto Encode é usada para codificar uma estrutura Go e salvá-la em formato JSON.
