package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Configuration de la difficulté
	// Pour un solveur naïf (sans propagation unitaire poussée) :
	// 20 vars = instantané
	// 40 vars = < 1 seconde
	// 50-60 vars = quelques secondes (le bon test pour toi)
	// 100 vars = potentiellement très long
	
	numVars := 60
	// Le ratio 4.26 est connu pour générer les problèmes les plus durs (Phase Transition)
	numClauses := int(float64(numVars) * 4.3) 

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("p cnf %d %d\n", numVars, numClauses)

	for i := 0; i < numClauses; i++ {
		clause := make([]int, 0, 3)
		used := make(map[int]bool)

		// Générer 3 littéraux distincts (3-SAT)
		for len(clause) < 3 {
			lit := rand.Intn(numVars) + 1
			
			// 50% de chance d'être négatif
			if rand.Intn(2) == 0 {
				lit = -lit
			}

			// On évite d'avoir x et -x ou deux fois x dans la même clause
			if !used[abs(lit)] {
				used[abs(lit)] = true
				clause = append(clause, lit)
			}
		}
		fmt.Printf("%d %d %d 0\n", clause[0], clause[1], clause[2])
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
