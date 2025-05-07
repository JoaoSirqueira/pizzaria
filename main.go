package main

import (
	 "github.com/gin-gonic/gin"
)

// type Pizza struct {
// 	ID int
// 	nome string
// 	preco float64
// }

func main() {
	router := gin.Default()
	// Criando uma rota
	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": "Toscana, Calabresa",
		})
	})
	
	// var pizzas = []Pizza{
	// 	{ID: 1, nome: "Toscana", preco: 49.5},
	// 	{ID: 2, nome: "Marguerita", preco: 79.5},
	// 	{ID: 3, nome: "Atum com Queijo", preco: 69.5},
	// }
	// fmt.Println(pizzas)
	router.Run() // Executar rota
}

// executar a operação: go run . 
// trazer dependências e pacotes externos: go get

// dependência do gin:
// go get github.com/gin-gonic/gin