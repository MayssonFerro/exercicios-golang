// 6 - que faça a ordenação de uma sequencia numérica

package main

import (
	"fmt"
	"sort"
)

func main() {
	var x int
	nums := make([]int, 0)

	fmt.Println("Digite números inteiros para ordenar. Digite 0 para terminar (0 não será incluído).")
	for i := 1; ; i++ {
		fmt.Printf("Digite o %dº número: ", i)
		if _, err := fmt.Scan(&x); err != nil {
			fmt.Println("Entrada inválida. Encerrando.")
			return
		}
		if x == 0 {
			break
		}
		nums = append(nums, x)
	}

	if len(nums) == 0 {
		fmt.Println("Nenhum número para ordenar.")
		return
	}

	sort.Ints(nums)

	fmt.Println("Sequência ordenada:")
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
