// 15 - calcule a área de um retângulo

package main

import "fmt"

func area(base, altura float64) float64 { return base * altura }

func main() {
	var b, h float64
	fmt.Print("Base: ")
	fmt.Scan(&b)
	fmt.Print("Altura: ")
	fmt.Scan(&h)
	fmt.Println(area(b, h))
}
