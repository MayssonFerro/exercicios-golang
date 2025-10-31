// 25 - que calcule a média de 2 ou mais números

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func media(nums []float64) float64 {
	s := 0.0
	for _, v := range nums {
		s += v
	}
	return s / float64(len(nums))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite números separados por espaço: ")
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	nums := make([]float64, 0, len(parts))
	for _, p := range parts {
		if v, err := strconv.ParseFloat(p, 64); err == nil {
			nums = append(nums, v)
		}
	}
	if len(nums) == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(media(nums))
}
