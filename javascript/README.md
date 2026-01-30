#  Flip 7 – Jeu en JavaScript (Node.js)

Ce projet est une implémentation en **mode texte** du jeu de cartes **Flip 7**, réalisée en JavaScript pour Node.js.

Le jeu permet à **plusieurs joueurs humains** de jouer à tour de rôle en partageant le clavier.  
Toutes les actions des joueurs sont **enregistrées dans un fichier de log**.

---

##  Prérequis

- Node.js (version récente)
- Un terminal (macOS, Linux ou Windows)

---

##  Lancer le jeu

Dans le dossier du projet :

```bash
node src/index.js
```
Le programme demande :
Le nombre de joueurs
Le nom de chaque joueur
À chaque tour, si le joueur souhaite tirer une carte ou s’arrêter
🎮 Règles du jeu (version implémentée)
Le paquet contient des cartes numérotées de 1 à 12
Il y a :
1 carte « 1 »
2 cartes « 2 »
3 cartes « 3 »
...
12 cartes « 12 »
À son tour, un joueur peut :
Tirer une carte
Ou s’arrêter volontairement
Si un joueur tire deux fois la même valeur, il perd immédiatement
Si un joueur atteint 7 cartes différentes, il gagne la manche
Le score final correspond à la somme des cartes du joueur

