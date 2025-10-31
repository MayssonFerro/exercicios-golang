// 11 - que recebe a data de nascimento e retorne o dia da semana

package main

import (
	"fmt"
	"time"
)

func diaSemana(data string) (string, error) {
	layouts := []string{"02/01/2006", "2006-01-02", "2/1/2006", time.RFC3339}
	var t time.Time
	var err error
	for _, l := range layouts {
		t, err = time.Parse(l, data)
		if err == nil {
			return t.Weekday().String(), nil
		}
	}
	return "", fmt.Errorf("formato de data inválido")
}

func main() {
	var s string
	fmt.Print("Digite a data de nascimento (DD/MM/AAAA): ")
	fmt.Scan(&s)
	dia, err := diaSemana(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dia)
}
