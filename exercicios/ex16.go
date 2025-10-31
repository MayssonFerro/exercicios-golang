// 16 - que calcule conversão entre unidades de temperatura

package main

import "fmt"

func cToF(c float64) float64 { return c*9/5 + 32 }
func fToC(f float64) float64 { return (f - 32) * 5 / 9 }

func main() {
	var t float64
	var u string
	fmt.Print("Digite valor e unidade (C/F): ")
	fmt.Scan(&t, &u)
	if u == "C" || u == "c" {
		fmt.Println(cToF(t))
		return
	}
	if u == "F" || u == "f" {
		fmt.Println(fToC(t))
		return
	}
	fmt.Println("Unidade inválida")
}
