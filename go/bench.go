package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func run(cmd string, args ...string) time.Duration {
	start := time.Now()
	exec.Command(cmd, args...).Run()
	return time.Since(start)
}

func main() {
	variables := []int{10, 20, 30, 40, 50}
	runs := 20

	fmt.Println("vars,seq_ms,par_ms")

	for _, v := range variables {
		var seqTotal time.Duration
		var parTotal time.Duration

		for i := 0; i < runs; i++ {

			// Génération de la formule
			exec.Command("go", "run", "generator.go", strconv.Itoa(v)).Run()

			// Solveur séquentiel
			seqTotal += run("go", "run", "main.go", "-mode=seq")

			// Solveur parallèle
			parTotal += run("go", "run", "main.go", "-mode=par")
		}

		seqAvg := float64(seqTotal.Milliseconds()) / float64(runs)
		parAvg := float64(parTotal.Milliseconds()) / float64(runs)

		fmt.Printf("%d,%.2f,%.2f\n", v, seqAvg, parAvg)
	}
}
