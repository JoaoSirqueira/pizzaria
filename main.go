package main

import (
	 "github.com/gin-gonic/gin"
)
type Pizza struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	router := gin.Default()
	// Criando uma rota
	router.GET("/pizzas", getPizzas)
	
	router.Run() // Executar rota
}

func getPizzas(c *gin.Context) {
	var pizzas = []Pizza{
		{ID: 1, Nome: "Toscana", Preco: 49.5},
		{ID: 2, Nome: "Marguerita", Preco: 79.5},
		{ID: 3, Nome: "Atum com Queijo", Preco: 69.5},
	}

	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
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