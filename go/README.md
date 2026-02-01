# Solveur SAT en Go (Séquentiel vs Parallèle)

Ce projet implémente un solveur SAT (Boolean Satisfiability Problem) basé sur l'algorithme **DPLL** (Davis-Putnam-Logemann-Loveland).

L'objectif principal est de comparer les performances entre :
1.  Une approche **séquentielle** (récursive classique).
2.  Une approche **parallèle** utilisant les Goroutines et le pattern Worker Pool de Go.

## Structure du projet

* `main.go` : Le code source du solveur (contient l'implémentation DPLL, le parseur DIMACS et le comparateur de performance).
* `generator.go` : Un utilitaire pour générer des problèmes SAT aléatoires et difficiles (Hard 3-SAT).
* `*.cnf` : Fichiers d'entrée au format DIMACS (formule logique).

## Prérequis

* [Go](https://go.dev/dl/) (version 1.18 ou supérieure recommandée).

## Utilisation du Générateur

Pour tester efficacement le parallélisme, il faut des problèmes suffisamment complexes. Le fichier `generator.go` crée des formules aléatoires.

### 1. Configuration
Ouvrez `generator.go` et modifiez la variable `numVars` pour ajuster la difficulté :

```go
// Dans generator.go
numVars := 60 // 20 = Instantané, 60 = Moyen, 100+ = Difficile
```
## 2. Génération
Le générateur affiche le résultat dans la sortie standard. Utilisez une redirection pour créer un fichier .cnf

```bash
# Générer un fichier de test
go run generator.go > test_complexe.cnf
```
Le générateur utilise un ratio clauses/variables de ~4.3, point critique connu pour créer les problèmes les plus difficiles à résoudre (transition de phase).

## Utilisation du Solveur
Le programme main.go prend en argument un fichier au format DIMACS (.cnf). Il exécute automatiquement les deux versions (séquentielle et parallèle) pour comparer les résultats.

```bash
go run main.go test_complexe.cnf
```
Exemple de sortie attendue :

```
Lecture du fichier : test_complexe.cnf
Nombre de clauses : 258

--- Test séquentiel ---
SATISFIABLE (séquentiel)
Temps séquentiel : 1.452000 s

--- Test parallèle ---
Cœurs CPU : 12
Profondeur parallèle automatique : 4
SATISFIABLE (parallel)
Temps parallèle : 0.484000 s

Accélération : ×3.00
```
## Analyse des performances

Il est crucial de comprendre que le parallélisme n'est pas toujours synonyme de vitesse :
1. Petits fichiers (< 30 variables) :

    * Le code séquentiel sera souvent plus rapide.

    * Raison : Le coût de création des Goroutines, de l'allocation mémoire et de la synchronisation (atomic/channels) est plus élevé que le calcul lui-même.

2. Gros fichiers (> 50 variables) :

    * Le code parallèle devient rentable.

    * Raison : La charge de travail par branche de l'arbre DPLL est assez lourde pour occuper efficacement tous les cœurs du processeur malgré le coût de gestion.


## Format de fichier supporté (DIMACS)
Le solveur accepte le format standard utilisé dans les compétitions SAT :

- Les lignes commençant par c sont des commentaires.
- La ligne p cnf [nbVars] [nbClauses] définit l'en-tête.
- Chaque ligne suivante est une clause se terminant par 0.


