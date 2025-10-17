// 13 - que retorne a moda de uma sequencia numérica

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func moda(nums []int) int {
	freq := make(map[int]int)
	best, bestCount := nums[0], 0
	for _, n := range nums {
		freq[n]++
		if freq[n] > bestCount {
			best = n
			bestCount = freq[n]
		}
	}
	return best
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite números separados por espaço: ")
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		if v, err := strconv.Atoi(p); err == nil {
			nums = append(nums, v)
		}
	}
	if len(nums) == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(moda(nums))
}
