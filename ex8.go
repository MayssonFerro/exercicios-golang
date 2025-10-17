// 8 - que implemente uma arvore de decisão binaria

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Node representa um nó da árvore de decisão
type Node struct {
	Question string
	Yes      *Node
	No       *Node
	Result   string
}

// Decide interativamente perguntando ao usuário
func (n *Node) DecideInteractive(reader *bufio.Reader) string {
	// Se for folha, retorna resultado
	if n.Result != "" {
		return n.Result
	}

	// Faz a pergunta
	fmt.Printf("%s (sim/nao): ", n.Question)
	resposta, _ := reader.ReadString('\n')
	resposta = strings.TrimSpace(strings.ToLower(resposta))

	// Trata resposta
	if resposta == "sim" || resposta == "s" || resposta == "y" || resposta == "yes" {
		return n.Yes.DecideInteractive(reader)
	} else if resposta == "nao" || resposta == "n" || resposta == "no" {
		return n.No.DecideInteractive(reader)
	} else {
		fmt.Println("❗ Resposta inválida, digite 'sim' ou 'nao'.")
		return n.DecideInteractive(reader) // repete a pergunta
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Folhas
	stayHome := &Node{Result: "Ficar em casa"}
	goOut := &Node{Result: "Sair"}

	// Nó intermediário
	isHot := &Node{
		Question: "Está quente?",
		Yes:      stayHome,
		No:       goOut,
	}

	// Raiz
	isRaining := &Node{
		Question: "Está chovendo?",
		Yes:      stayHome,
		No:       isHot,
	}

	fmt.Println("Sistema de decisão:")
	decisao := isRaining.DecideInteractive(reader)
	fmt.Println("\nDecisão final:", decisao)
}
