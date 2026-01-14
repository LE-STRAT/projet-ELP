package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
)

////////////////////////////
// Types
////////////////////////////

type Literal int
type Clause []Literal
type Formula []Clause

type Task struct {
	formula Formula
	depth   int
}

////////////////////////////
// Utilitaires
////////////////////////////

func abs(x Literal) Literal {
	if x < 0 {
		return -x
	}
	return x
}

func computeParallelDepth(numCPU int) int {
	depth := 0
	tasks := 1
	for tasks < numCPU {
		tasks *= 2
		depth++
	}
	return depth
}

////////////////////////////
// Simplification
////////////////////////////

func simplify(formula Formula, variable Literal, value bool) Formula {
	newFormula := Formula{}

	for _, clause := range formula {
		newClause := Clause{}
		satisfied := false

		for _, lit := range clause {

			if lit == variable && value {
				satisfied = true
				break
			}
			if lit == -variable && !value {
				satisfied = true
				break
			}

			if lit != variable && lit != -variable {
				newClause = append(newClause, lit)
			}
		}

		if !satisfied {
			newFormula = append(newFormula, newClause)
		}
	}

	return newFormula
}

////////////////////////////
// DPLL séquentiel
////////////////////////////

func dpll(formula Formula) bool {
	if len(formula) == 0 {
		return true
	}

	for _, clause := range formula {
		if len(clause) == 0 {
			return false
		}
	}

	variable := abs(formula[0][0])

	for _, value := range []bool{true, false} {
		newFormula := simplify(formula, variable, value)
		if dpll(newFormula) {
			return true
		}
	}

	return false
}

////////////////////////////
// Worker pool
////////////////////////////

func worker(id int, tasks <-chan Task, found *atomic.Bool, result chan<- bool, maxDepth int) {
	for task := range tasks {

		if found.Load() {
			return
		}

		if len(task.formula) == 0 {
			found.Store(true)
			result <- true
			return
		}

		for _, clause := range task.formula {
			if len(clause) == 0 {
				goto nextTask
			}
		}

		variable := abs(task.formula[0][0])

		for _, value := range []bool{true, false} {

			if found.Load() {
				return
			}

			newFormula := simplify(task.formula, variable, value)

			if task.depth < maxDepth {
				// Créer une nouvelle tâche
				tasksChan <- Task{
					formula: newFormula,
					depth:   task.depth + 1,
				}
			} else {
				// Séquentiel
				if dpll(newFormula) {
					found.Store(true)
					result <- true
					return
				}
			}
		}

	nextTask:
		continue
	}
}

////////////////////////////
// Solveur parallèle
////////////////////////////

var tasksChan chan Task

func parallelSolve(formula Formula) bool {

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	maxDepth := computeParallelDepth(numCPU)

	fmt.Println("Cœurs CPU :", numCPU)
	fmt.Println("Profondeur parallèle automatique :", maxDepth)

	tasksChan = make(chan Task, 10000)
	result := make(chan bool)
	var found atomic.Bool

	for i := 0; i < numCPU; i++ {
		go worker(i, tasksChan, &found, result, maxDepth)
	}

	// Tâche initiale
	tasksChan <- Task{
		formula: formula,
		depth:   0,
	}

	for {
		select {
		case r := <-result:
			if r {
				close(tasksChan)
				return true
			}
			if found.Load() {
				close(tasksChan)
				return true
			}
		}
	}
}

////////////////////////////
// Lecture DIMACS
////////////////////////////

func parseDIMACS(filename string) (Formula, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	formula := Formula{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "c") {
			continue
		}

		if strings.HasPrefix(line, "p") {
			continue
		}

		fields := strings.Fields(line)
		clause := Clause{}

		for _, f := range fields {
			val, err := strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
			if val == 0 {
				break
			}
			clause = append(clause, Literal(val))
		}

		if len(clause) > 0 {
			formula = append(formula, clause)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return formula, nil
}

////////////////////////////
// Main
////////////////////////////

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go fichier.cnf")
		return
	}

	filename := os.Args[1]

	fmt.Println("Lecture du fichier :", filename)

	formula, err := parseDIMACS(filename)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("Nombre de clauses :", len(formula))

	if parallelSolve(formula) {
		fmt.Println("SATISFIABLE (parallel)")
	} else {
		fmt.Println("UNSATISFIABLE (parallel)")
	}
}
