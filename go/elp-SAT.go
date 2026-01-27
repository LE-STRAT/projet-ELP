package main

import "fmt"

type Clause []int
type Formula []Clause

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func simplify(formula Formula, variable int, value bool) Formula {
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

func main() {
    formula := Formula{
        Clause{1, -2},
        Clause{2},
    }

    if dpll(formula) {
        fmt.Println("SATISFIABLE")
    } else {
        fmt.Println("UNSATISFIABLE")
    }
}
