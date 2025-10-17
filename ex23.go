// 23 - que calcule IMC padrão Brasil

package main

import "fmt"

func imc(peso, altura float64) float64 { return peso / (altura * altura) }

func classificaIMC(imc float64) string {
	switch {
	case imc < 18.5:
		return "Abaixo do peso"
	case imc < 25.0:
		return "Peso normal"
	case imc < 30.0:
		return "Sobrepeso"
	case imc < 35.0:
		return "Obesidade grau I"
	case imc < 40.0:
		return "Obesidade grau II"
	default:
		return "Obesidade grau III"
	}
}

func main() {
	var peso, altura float64
	fmt.Print("Peso (kg): ")
	fmt.Scan(&peso)
	fmt.Print("Altura (m): ")
	fmt.Scan(&altura)
	if peso <= 0 || altura <= 0 {
		fmt.Println("Peso e altura devem ser maiores que zero")
		return
	}
	// se altura provavelmente foi informada em centímetros (ex: 181), converte para metros
	if altura > 10 {
		altura = altura / 100.0
	}
	val := imc(peso, altura)
	fmt.Printf("IMC: %.2f - %s\n", val, classificaIMC(val))
}
