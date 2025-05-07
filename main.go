package main

import (
	"pizzaria/models"
	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Nome: "Toscana", Preco: 49.5},
	{ID: 2, Nome: "Marguerita", Preco: 79.5},
	{ID: 3, Nome: "Atum com Queijo", Preco: 69.5},
}

func main() {
	router := gin.Default()
	// Criando uma rota
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.Run() // Executar rota
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}
	pizzas = append(pizzas, newPizza)
}


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