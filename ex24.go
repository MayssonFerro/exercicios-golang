// 24 - que calcule MMC

package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int { return a / gcd(a, b) * b }

func main() {
	var a, b int
	fmt.Print("Digite dois números: ")
	fmt.Scan(&a, &b)
	fmt.Println(lcm(a, b))
}
