package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	img, _ := os.Open("image.ppm")

	reader := bufio.NewReader(img)

	var formato string
	var linhas, colunas, maxVal int

	fmt.Fscan(reader, &formato)
	fmt.Fscan(reader, &linhas, &colunas)
	fmt.Fscan(reader, &maxVal)

	fmt.Printf("Formato: %s\n", formato)
	fmt.Printf("Dimensões: %dx%d\n", linhas, colunas)
	fmt.Printf("Valor máximo: %d\n", maxVal)
	fmt.Printf("Pixels:\n")

	var r, g, b int

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			fmt.Fscan(reader, &r, &g, &b)
			_ = r
			_ = g
			_ = b
			fmt.Printf("%d %d %d\n", r, g, b)
		}
	}
}
