// 10 - que resolva a brincadeira de peças da torre de Hanói de três discos

package main

import (
	"fmt"
	"strings"
)

func hanoi(n int, origem, destino, auxiliar rune, pegs map[rune][]int) {
	if n <= 0 {
		return
	}
	if n == 1 {
		disk := pegs[origem][len(pegs[origem])-1]
		pegs[origem] = pegs[origem][:len(pegs[origem])-1]
		pegs[destino] = append(pegs[destino], disk)
		fmt.Printf("Mover disco %d de %c para %c\n", disk, origem, destino)
		return
	}
	hanoi(n-1, origem, auxiliar, destino, pegs)
	hanoi(1, origem, destino, auxiliar, pegs)
	hanoi(n-1, auxiliar, destino, origem, pegs)
}

func formatPeg(p []int) string {
	if len(p) == 0 {
		return "[]"
	}
	// queremos mostrar do maior para o menor (base -> topo)
	parts := make([]string, len(p))
	for i := len(p) - 1; i >= 0; i-- {
		parts[len(p)-1-i] = fmt.Sprintf("%d", p[i])
	}
	return "[" + strings.Join(parts, " ") + "]"
}

func main() {
	n := 3
	pegs := map[rune][]int{
		'A': make([]int, 0, n),
		'B': make([]int, 0, n),
		'C': make([]int, 0, n),
	}
	for d := n; d >= 1; d-- {
		pegs['A'] = append(pegs['A'], d)
	}

	fmt.Printf("Ordem inicial: A:%s B:%s C:%s\n", formatPeg(pegs['A']), formatPeg(pegs['B']), formatPeg(pegs['C']))

	hanoi(n, 'A', 'C', 'B', pegs)
}
