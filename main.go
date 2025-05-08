package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()
	router := gin.Default()
	// Criando uma rota
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasByID)
	router.DELETE("/pizzas/:id", deletePizzasByID) // Deletar uma pizza
	router.PUT("/pizzas/:id", updatePizzasByID)// Editar ou atualizar uma pizza
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
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(201, newPizza)
}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}
	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza not found"})
}

func loadPizzas() {
	file, err := os.Open("data/pizza.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close() // Instrução para fechar.
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func savePizza() {
	file, err := os.Create("data/pizza.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close() // Instrução para fechar.
	encoder := json.NewEncoder(file) // Ler arquivo e transformar em arquivos json
	if err := encoder.Encode(pizzas); err != nil{
		fmt.Println("Error encoding JSON:", err)
	}
}

func deletePizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	 for i, p := range pizzas {
		if p.ID == id {
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			savePizza()
			c.JSON(200, gin.H{"message":"pizza deleted"})
			return
		}
	 }
	c.JSON(404, gin.H{"message":"pizza not found"})
}

func updatePizzasByID(c *gin.Context) {
	c.JSON(200, gin.H{"method":"put"})
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
