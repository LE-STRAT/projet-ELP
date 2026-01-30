import Game from "./game.js";
import readline from "readline";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

function ask(question) {
  return new Promise(resolve => rl.question(question, resolve));
}

const count = Number(await ask("Nombre de joueurs : "));

const names = [];
for (let i = 0; i < count; i++) {
  const name = await ask(`Nom du joueur ${i + 1} : `);
  names.push(name);
}

rl.close();

const game = new Game(names);
await game.play();
