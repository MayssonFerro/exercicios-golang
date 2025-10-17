// 19 - que conte a quantidade de vogais e consoantes presentes em um texto

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func conta(s string) (int, int) {
	vogais := "aeiou"
	v, c := 0, 0
	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) {
			continue
		}
		if strings.ContainsRune(vogais, r) {
			v++
		} else {
			c++
		}
	}
	return v, c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite texto: ")
	s, _ := reader.ReadString('\n')
	v, c := conta(s)
	fmt.Println(v, "vogais e", c, "consoantes")
}
