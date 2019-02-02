package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json

	p1 := Product{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônicos"}}
	piJSON, _ := json.Marshal(p1)

	fmt.Println(string(piJSON))

	var p2 Product

	jsonString := `{"id":1,"name":"Caneta","price":9.9,"tags":["Papelaria","Importado"]}`

	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
