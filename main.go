package main

import (
    "bufio"
    "fmt"
    "os"
    "runtime"
    "strconv"
    "strings"
    "sync/atomic"
    "time"
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

var tasksChan chan Task
var activeTasks atomic.Int64

func worker(id int, tasks <-chan Task, found *atomic.Bool, result chan<- bool, maxDepth int, doneSent *atomic.Bool) {
    for {
        if found.Load() {
            return
        }

        task, ok := <-tasks
        if !ok {
            return
        }

        // Cas terminaux
        if len(task.formula) == 0 {
            if found.CompareAndSwap(false, true) {
                result <- true
            }
            activeTasks.Add(-1)
            return
        }

        hasEmptyClause := false
        for _, clause := range task.formula {
            if len(clause) == 0 {
                hasEmptyClause = true
                break
            }
        }

        if hasEmptyClause {
            activeTasks.Add(-1)
            if activeTasks.Load() == 0 && !found.Load() && doneSent.CompareAndSwap(false, true) {
                result <- false
            }
            continue
        }

        variable := abs(task.formula[0][0])

        for _, value := range []bool{true, false} {

            if found.Load() {
                activeTasks.Add(-1)
                return
            }

            newFormula := simplify(task.formula, variable, value)

            if task.depth < maxDepth {
                // Créer une nouvelle tâche
                activeTasks.Add(1)
                tasksChan <- Task{
                    formula: newFormula,
                    depth:   task.depth + 1,
                }
            } else {
                // Séquentiel
                if dpll(newFormula) {
                    if found.CompareAndSwap(false, true) {
                        result <- true
                    }
                    activeTasks.Add(-1)
                    return
                }
            }
        }

        // Fin du traitement de cette tâche
        activeTasks.Add(-1)
        if activeTasks.Load() == 0 && !found.Load() && doneSent.CompareAndSwap(false, true) {
            result <- false
            return
        }
    }
}

////////////////////////////
// Solveur parallèle
////////////////////////////

func parallelSolve(formula Formula) bool {

    numCPU := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPU)

    maxDepth := computeParallelDepth(numCPU)

    fmt.Println("Cœurs CPU :", numCPU)
    fmt.Println("Profondeur parallèle automatique :", maxDepth)

    tasksChan = make(chan Task, 10000)
    result := make(chan bool, 1)
    var found atomic.Bool
    var doneSent atomic.Bool

    activeTasks.Store(1)

    for i := 0; i < numCPU; i++ {
        go worker(i, tasksChan, &found, result, maxDepth, &doneSent)
    }

    // Tâche initiale
    tasksChan <- Task{
        formula: formula,
        depth:   0,
    }

    // On attend le premier résultat (SAT ou UNSAT)
    r := <-result
    return r
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

    ////////////////////////////
    // Test séquentiel
    ////////////////////////////

    fmt.Println("\n--- Test séquentiel ---")
    startSeq := time.Now()
    seqResult := dpll(formula)
    durationSeq := time.Since(startSeq)

    if seqResult {
        fmt.Println("SATISFIABLE (séquentiel)")
    } else {
        fmt.Println("UNSATISFIABLE (séquentiel)")
    }
    fmt.Printf("Temps séquentiel : %.6f s\n", durationSeq.Seconds())

    ////////////////////////////
    // Test parallèle
    ////////////////////////////

    fmt.Println("\n--- Test parallèle ---")
    startPar := time.Now()
    parResult := parallelSolve(formula)
    durationPar := time.Since(startPar)

    if parResult {
        fmt.Println("SATISFIABLE (parallel)")
    } else {
        fmt.Println("UNSATISFIABLE (parallel)")
    }
    fmt.Printf("Temps parallèle : %.6f s\n", durationPar.Seconds())

    ////////////////////////////
    // Comparaison
    ////////////////////////////

    speedup := durationSeq.Seconds() / durationPar.Seconds()
    fmt.Printf("\nAccélération : ×%.2f\n", speedup)
}