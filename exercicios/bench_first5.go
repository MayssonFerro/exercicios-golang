package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func runCmd(cmdName string, args []string, input string) (time.Duration, string, error) {
	cmd := exec.Command(cmdName, args...)
	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}
	start := time.Now()
	out, err := cmd.CombinedOutput()
	elapsed := time.Since(start)
	return elapsed, string(out), err
}

func main() {
	exs := []struct {
		file  string
		input string
	}{
		{"ex1.go", "5\n3\n"},
		{"ex2.go", "10\n2\n"},
		{"ex3.go", "7\n"},
		{"ex4.go", "4\n"},
		{"ex5.go", "17\n"},
	}

	fmt.Println("Benchmark dos 5 primeiros exercícios (5 execuções cada):")
	for _, ex := range exs {
		var times []time.Duration
		var lastOut string
		for i := 0; i < 5; i++ {
			elapsed, out, err := runCmd("go", []string{"run", ex.file}, ex.input)
			if err != nil {
				fmt.Printf("Erro ao executar %s: %v\n", ex.file, err)
			}
			times = append(times, elapsed)
			lastOut = out
		}
		// compute average
		var total time.Duration
		for _, t := range times {
			total += t
		}
		avg := total / time.Duration(len(times))
		fmt.Printf("%s - média: %v (runs: ", ex.file, avg)
		for i, t := range times {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(t)
		}
		fmt.Println(")")
		fmt.Printf("Última saída:\n%s\n", lastOut)
	}
}
