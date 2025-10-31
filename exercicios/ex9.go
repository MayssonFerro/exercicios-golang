// 9 - imprime o valor e endereço de uma variável
package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um número inteiro: ")
	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("Erro na leitura:", err)
		return
	}

	fmt.Println("Valor de x:", x)
	fmt.Printf("Endereço de x: %p\n", &x)
}
