// 18 - que receba 3 matrizes de inteiros sem sinal de 8 bits e grave um arquivo .jpg

// Exemplo de funcionamento:
// Entrada:
// Largura altura: 2 2
// 255 0 0 255
// 0 255 0 0
// 0 0 255 255
// Saída: arquivo out.jpg criado

package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	var w, h int
	fmt.Print("Largura altura: ")
	fmt.Scan(&w, &h)
	total := w * h
	r := make([]uint8, total)
	g := make([]uint8, total)
	b := make([]uint8, total)
	fmt.Printf("Digite %d bytes para R, G e B (decimal, separados por espaço ou newline):\n", total)
	for i := 0; i < total; i++ {
		var v int
		fmt.Scan(&v)
		r[i] = uint8(v)
	}
	for i := 0; i < total; i++ {
		var v int
		fmt.Scan(&v)
		g[i] = uint8(v)
	}
	for i := 0; i < total; i++ {
		var v int
		fmt.Scan(&v)
		b[i] = uint8(v)
	}
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	idx := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{r[idx], g[idx], b[idx], 255})
			idx++
		}
	}
	f, err := os.Create("out.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var opt jpeg.Options
	opt.Quality = 80
	_ = binary.Write(f, binary.LittleEndian, []byte{})
	jpeg.Encode(f, img, &opt)
	fmt.Println("out.jpg gerado")
}
