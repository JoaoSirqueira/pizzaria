package data

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("dados/pizza.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close() // Instrução para fechar.
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func SavePizza() {
	file, err := os.Create("dados/pizza.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close() // Instrução para fechar.
	encoder := json.NewEncoder(file) // Ler arquivo e transformar em arquivos json
	if err := encoder.Encode(Pizzas); err != nil{
		fmt.Println("Error encoding JSON:", err)
	}
}