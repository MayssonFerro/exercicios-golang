// 4 - verificar se um número é par, positivo ou negativo

package main

import "fmt"

func ehPar(num int) bool {
	return num%2 == 0
}

func classificarNumero(num int) string {
	if num > 0 {
		return "positivo"
	} else if num < 0 {
		return "negativo"
	} else {
		return "zero"
	}
}

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if ehPar(numero) {
		fmt.Printf("O número %d é PAR\n", numero)
	} else {
		fmt.Printf("O número %d é ÍMPAR\n", numero)
	}

	classificacao := classificarNumero(numero)
	fmt.Printf("O número %d é %s\n", numero, classificacao)
}
