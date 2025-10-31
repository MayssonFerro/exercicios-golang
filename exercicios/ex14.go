// 14 - verifique se uma sequencia de caracteres formam um palíndromo

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func palindromo(s string) bool {
	rs := make([]rune, 0, len(s))
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			rs = append(rs, r)
		}
	}
	i, j := 0, len(rs)-1
	for i < j {
		if rs[i] != rs[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma sequência: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println(palindromo(s))
}
