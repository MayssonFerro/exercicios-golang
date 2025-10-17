// 5 - verificar se um número é primo ou não

package main

import "fmt"

func ehPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if ehPrimo(numero) {
		fmt.Printf("%d é primo\n", numero)
	} else {
		fmt.Printf("%d não é primo\n", numero)
	}
}
