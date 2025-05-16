package main

import "fmt"

func main() {
	carros := map[string]int{"Fusca":3000, "Civic":2000, "Palio":1500}
	
	for nome, valor := range carros {
		fmt.Printf("O carro %s custa %d\n", nome, valor)
		
	}
	
}