// 2 - divisão de dois números inteiros

package main

import "fmt"

func divisao(a, b int) int {
	return a / b
}

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	resultado := divisao(num1, num2)

	fmt.Printf("%d / %d = %d\n", num1, num2, resultado)
}
