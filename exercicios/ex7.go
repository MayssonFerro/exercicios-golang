// 7 - que faça a ordenação de caracteres em ordem ascendente

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Digite uma linha de caracteres:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && len(input) == 0 {
		return
	}

	input = strings.TrimRight(input, "\r\n")

	runes := []rune(input)

	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

	parts := make([]string, len(runes))
	for i, r := range runes {
		parts[i] = string(r)
	}

	fmt.Println("Caracteres ordenados:", strings.Join(parts, " "))
}
