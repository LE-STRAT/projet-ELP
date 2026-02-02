package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generator.go <numVars> [numClauses]")
		return
	}

	numVars, err := strconv.Atoi(os.Args[1])
	if err != nil || numVars <= 0 {
		fmt.Println("numVars doit être un entier positif")
		return
	}

	numClauses := numVars * 4
	if len(os.Args) >= 3 {
		c, err := strconv.Atoi(os.Args[2])
		if err == nil && c > 0 {
			numClauses = c
		}
	}

	f, err := os.Create("formula.cnf")
	if err != nil {
		fmt.Println("Erreur création fichier:", err)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "p cnf %d %d\n", numVars, numClauses)

	for i := 0; i < numClauses; i++ {
		clauseSize := rand.IntN(3) + 1
		for j := 0; j < clauseSize; j++ {
			v := rand.IntN(numVars) + 1
			if rand.IntN(2) == 0 {
				v = -v
			}
			fmt.Fprintf(f, "%d ", v)
		}
		fmt.Fprintf(f, "0\n")
	}
}
