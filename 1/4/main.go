package main

import "fmt"

func main() {
	nomes := [5]string{"David", "Paulo", "João", "Maria", "Ana"}
    // definindo o array com 5 posições
	nomes[0] = "David"
	nomes[1] = "Paulo"
	nomes[2] = "João"
	nomes[3] = "Maria"
	nomes[4] = "Ana"

	for i, v := range nomes {
		fmt.Printf("O nome %s está na posição %d\n", v, i)
	}
}