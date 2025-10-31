// 17 - que simule o jogo da adivinhação

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int
	for {
		fmt.Print("Adivinhe (1-100): ")
		fmt.Scan(&guess)
		if guess < target {
			fmt.Println("Maior")
		} else if guess > target {
			fmt.Println("Menor")
		} else {
			fmt.Println("Acertou")
			break
		}
	}
}
