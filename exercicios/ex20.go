// 20 - que encontre padrões de ocorrências de palavras em um texto

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ocorrencias(text, pattern string) int {
	re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(pattern))
	return len(re.FindAllStringIndex(text, -1))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o texto: ")
	txt, _ := reader.ReadString('\n')
	fmt.Print("Digite a palavra/padrão: ")
	pat, _ := reader.ReadString('\n')
	txt = strings.TrimSpace(txt)
	pat = strings.TrimSpace(pat)
	fmt.Println(ocorrencias(txt, pat))
}
