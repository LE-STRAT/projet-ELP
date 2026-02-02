# Solveur SAT en Go — Séquentiel vs Parallèle

Ce projet implémente un solveur SAT basé sur l’algorithme **DPLL** (Davis–Putnam–Logemann–Loveland) et compare deux approches :

- une version **séquentielle**
- une version **parallèle** utilisant les goroutines

Le projet inclut également :

- un **générateur de formules DIMACS**
- un **benchmark automatisé** produisant un CSV
- un **script Python** pour tracer les courbes de performance

---

## Structure du projet

| Fichier | Description |
|--------|-------------|
| `main.go` | Solveur SAT + mode benchmark (`-mode=seq` / `-mode=par`) |
| `generator.go` | Générateur de formules DIMACS aléatoires |
| `bench.go` | Benchmark automatisé (génère un CSV) |
| `plot.py` | Script Python pour tracer les courbes |
| `formula.cnf` | Dernière formule générée |
| `results.csv` | Résultats du benchmark |

---

## Prérequis

- Go **1.25+**
- Python 3
- matplotlib 

---

# Génération de formules SAT

Le générateur produit une formule DIMACS aléatoire dans `formula.cnf`

### Utilisation
```
go run generator.go  <numVars> [numClauses]
```

- `numVars` : nombre de variables (obligatoire)
- `numClauses` : optionnel (par défaut = `numVars * 4`)

Exemple :
```
go run generator.go  40
```

# Utilisation du solveur (mode normal)

Pour résoudre un fichier DIMACS :

go run main.go  fichier.cnf

Le solveur :

- charge la formule
- exécute la version séquentielle
- exécute la version parallèle
- affiche les temps et le speedup

---

# Mode benchmark (automatisé)

Le solveur possède un mode spécial utilisé par `bench.go` :
```
go run main.go  -mode=seq
go run main.go  -mode=par
```

Dans ce mode :

- le solveur lit automatiquement `formula.cnf`
- il exécute uniquement la version demandée
- il affiche **uniquement le temps en millisecondes**

---

# Benchmark automatisé

Le script `bench.go` :

- génère plusieurs formules pour différentes tailles
- exécute les solveurs séquentiel et parallèle
- calcule les moyennes
- produit un fichier CSV

### Lancer le benchmark

### ✔ Linux / macOS

go run bench.go  > results.csv

# Tracer les courbes

Le script Python lit `results.csv` et trace :

- temps séquentiel
- temps parallèle

### Lancer :
```
python3 plot.py
```
Une fenêtre matplotlib s’ouvre avec le graphe.

---

# Format DIMACS supporté

- `c ...` : commentaires  
- `p cnf <nbVars> <nbClauses>` : en-tête  
- chaque clause se termine par `0`

Exemple :
```
p cnf 3 2
1 -2 0
3 0
```

# Analyse des performances

### Petites formules (< 30 variables)

- Le séquentiel est souvent plus rapide.
- Le coût du parallélisme dépasse le coût du calcul.

### Formules plus grandes (> 50 variables)

- Le parallèle devient rentable.
- Le speedup dépend fortement de la structure de la formule.

---
