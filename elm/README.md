# TcTurtle Visualizer

Un visualiseur interactif pour le langage de programmation graphique TcTurtle.  
Dessinez des figures géométriques en contrôlant les mouvements d'une tortue virtuelle !


## Les fonctionnalités :
-  Un éditeur de code en direct : vous écrivez votre programme TcTurtle et visualisez le résultat en direct
-  Commandes simples tel que « Forward », « Left », « Right » et « Repeat »
- Gestion des erreurs : Messages si la syntaxe est incorrecte et guide pour commencer


## Les prérequis : 
Il vous faudra installer ELM (version 0.19+) et un terminal/invite de commandes


## L’installation :

Étape 1 : Accédez au dossier du projet

Étape 2 : Lancez « elm reactor »

Étape 3 : Ouvrez votre navigateur et allez à : http://localhost:8000/Main.elm


## Les commandes disponibles :
Pour avancer de N pixels : Forward N  
Pour tourner à gauche de N degrés : Left N  
Pour tourner à droite de N degrés : Left N  
Pour répéter N fois les instructions : Repeat N [...]


## Le Format :
- Toutes les instructions doivent être entre crochets []
- Les instructions doivent être séparées par des virgules
- Les espaces sont ignorés
- Les nombres doivent être des entiers


## Description des modules
### Main.elm
- Gère l'interface utilisateur
- Éditeur de code et affichage des erreurs
- Section d'aide intégrée
- Communication avec le parseur

### Drawing.elm
- Interprète les instructions de la tortue
- Gère la position (x, y) et l'orientation (angle)
- Génère les lignes à dessiner en SVG
- Affiche le résultat graphique

### ParserTcTurtle.elm
- Parse le code TcTurtle en utilisant la librairie Parser
- Valide la syntaxe
- Retourne les erreurs ou le programme compilé


## Les erreurs récurrentes : 
- Vérifiez que le code est entre crochets [ ]
- Ne pas séparer les instructions par des virgules
- Attention à l’orthographe des instructions

Si le dessin ne s'affiche pas :
- Vérifiez les crochets fermants
- Assurez-vous que tous les nombres sont des entiers
- Consultez la section d'aide pour voir un exemple valide


## Les languages utilisées :
- Elm 0.19 : pour le frontend
- HTML/CSS : pour l’interface utilisateur


## Motifs à tester :

—> Carré : [Forward 100, Right 90, Forward 100, Right 90, Forward 100, Right 90, Forward 100]
 
—> Triangle avec répétition : [Repeat 3 [Forward 100, Right 120]]

—> Étoile simple  : [Repeat 5 [Forward 100, Right 72]]

—> Spirale : [Repeat 10 [Forward 30, Right 36, Forward 60, Right 36]]

—> Escalier : [Forward 50, Right 90, Forward 50, Left 90, Forward 50, Right 90, Forward 50, Left 90, Forward 50]

—> Forme complexe : [Repeat 4 [Forward 100, Left 90, Forward 50, Right 90]]
