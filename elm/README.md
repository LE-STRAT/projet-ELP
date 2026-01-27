# 🐢 TcTurtle Visualizer

Un visualiseur interactif pour le langage de programmation graphique **TcTurtle**. Dessinez des figures géométriques en contrôlant les mouvements d'une tortue virtuelle !

## 🎯 Fonctionnalités

- ✅ **Éditeur de code en direct** : Écrivez votre programme TcTurtle et visualisez le résultat en temps réel
- ✅ **Syntaxe intuitive** : Commandes simples comme `Forward`, `Left`, `Right`, `Repeat`
- ✅ **Gestion d'erreurs** : Messages clairs en cas de syntaxe incorrecte
- ✅ **Aide intégrée** : Guide de syntaxe directement dans l'interface
- ✅ **Interface épurée** : Design moderne et ergonomique
- ✅ **Responsive** : Fonctionne sur différentes tailles d'écran

## 🚀 Démarrage rapide

### Prérequis

- [Elm](https://elm-lang.org/) installé (version 0.19+)
- Un terminal/invite de commandes

### Installation

1. Accédez au dossier du projet :
```bash
cd bidule\truc\elp
```

2. Lancez `elm reactor` :
```bash
elm reactor
```

3. Ouvrez votre navigateur et allez à :
```
http://localhost:8000/Main.elm
```

Voilà ! L'application est prête à l'emploi ! 🎉

## 📖 Syntaxe du langage TcTurtle

### Commandes disponibles

| Commande | Description | Exemple |
|----------|-------------|---------|
| `Forward N` | Avancer de N pixels | `Forward 100` |
| `Left N` | Tourner à gauche de N degrés | `Left 90` |
| `Right N` | Tourner à droite de N degrés | `Right 90` |
| `Repeat N [...]` | Répéter N fois les instructions | `Repeat 4 [Forward 100, Right 90]` |

### Format

- Toutes les instructions doivent être entre **crochets `[]`**
- Les instructions doivent être **séparées par des virgules `,`**
- Les espaces sont **ignorés**
- Les nombres doivent être des **entiers**

### Exemples

#### Carré
```
[Forward 100, Right 90, Forward 100, Right 90, Forward 100, Right 90, Forward 100]
```

#### Triangle avec répétition
```
[Repeat 3 [Forward 100, Right 120]]
```

#### Forme complexe
```
[Repeat 4 [Forward 100, Left 90, Forward 50, Right 90]]
```

## 🏗️ Architecture du projet

```
elm/
├── src/
│   ├── Main.elm              # Interface utilisateur principale
│   ├── Drawing.elm           # Moteur de dessin (SVG)
│   └── ParserTcTurtle.elm    # Parseur du langage TcTurtle
├── elm.json                  # Configuration Elm
└── README.md                 # Ce fichier
```

### Description des modules

**Main.elm**
- Gère l'interface utilisateur
- Éditeur de code et affichage des erreurs
- Section d'aide intégrée
- Communication avec le parseur

**Drawing.elm**
- Interprète les instructions de la tortue
- Gère la position (x, y) et l'orientation (angle)
- Génère les lignes à dessiner en SVG
- Affiche le résultat graphique

**ParserTcTurtle.elm**
- Parse le code TcTurtle en utilisant la librairie `Parser`
- Valide la syntaxe
- Retourne les erreurs ou le programme compilé

## 🎓 Comment ça marche ?

1. **Vous écrivez** du code TcTurtle dans l'éditeur
2. **Vous cliquez** sur le bouton "▶ Dessiner"
3. **Le parser** vérifie la syntaxe et crée les instructions
4. **Le moteur de dessin** exécute les instructions pas à pas :
   - Déplace la tortue
   - Trace les lignes
   - Tourne la tortue
5. **Le résultat** s'affiche en SVG sur la droite

## 🐛 Résolution des problèmes

### "Erreur de syntaxe ! Vérifiez votre code."
- Vérifiez que le code est entre crochets `[ ]`
- Assurez-vous que les instructions sont séparées par des **virgules**
- Vérifiez l'orthographe : `Forward` (pas `forward`)

### Le dessin ne s'affiche pas
- Vérifiez les crochets fermants
- Assurez-vous que tous les nombres sont des entiers
- Consultez la section d'aide pour voir un exemple valide

### Les lignes sortent de l'écran
- C'est normal ! La tortue démarre au centre (250, 250)
- Essayez avec des valeurs plus petites (`Forward 50`)

## 🛠️ Technologies utilisées

- **Elm 0.19** : Langage fonctionnel pour le frontend
- **Parser** : Librairie Elm pour le parsing robuste
- **SVG** : Rendu graphique vectoriel
- **HTML/CSS** : Interface utilisateur

## 📝 Exemples d'utilisation

### Étoile simple
```
[Repeat 5 [Forward 100, Right 72]]
```

### Spirale
```
[Repeat 10 [Forward 30, Right 36, Forward 60, Right 36]]
```

### Escalier
```
[Forward 50, Right 90, Forward 50, Left 90, Forward 50, Right 90, Forward 50, Left 90, Forward 50]
```

## 👥 Contribution

Ce projet a été développé dans le cadre du cours ELP à l'INSA.

## 📄 Licence

Libre d'utilisation à titre éducatif.

---

**Bon dessin ! 🎨**
