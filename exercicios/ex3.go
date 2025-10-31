// 3 - um algoritmo que recebe um número inteiro e imprime o antecessor e sucessor

package main

import "fmt"

func antecessor(a int) int {
	return a - 1
}

func sucessor(a int) int {
	return a + 1
}

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	fmt.Printf("O antecessor de %d é %d\n", num, antecessor(num))
	fmt.Printf("O sucessor de %d é %d\n", num, sucessor(num))
}
