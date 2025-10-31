// 12 - que retorne um booleano e a igualdade entre dois números

package main

import "fmt"

func igual(a, b int) bool { return a == b }

func main() {
	var a, b int
	fmt.Print("Digite dois números: ")
	fmt.Scan(&a, &b)
	fmt.Println(igual(a, b))
}
