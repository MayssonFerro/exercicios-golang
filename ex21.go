// 21 - que calcule o fatorial de um número n

package main

import "fmt"

func fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}

func main() {
	var n int
	fmt.Print("Digite n: ")
	fmt.Scan(&n)
	fmt.Println(fatorial(n))
}
