import Deck from "./deck.js";
import Player from "./player.js";
import Input from "./input.js";
import Logger from "./logger.js";

export default class Game {
	constructor(playerName) {
		this.deck = new Deck();
		this.players = playerName.map(name => new Player(name));
		this.input = new Input();
		this.logger = new Logger();
	}

	async play() {
		console.log(" Début de la partie Flip 7\n");
		this.logger.log("Début de la partie");

		for(const player of this.players) {
			await this.playPlayerTurn(player);
			this.logger.log(`Fin du tour de ${player.name}`)
			console.log("\n----------------------\n");
		}

		this.showScores();
		this.input.close();
		console.log(" Fin de la manche");
	}

	async playPlayerTurn(player) {
		console.log(`Tour de ${player.name}`);
		this.logger.log(`Tour de ${player.name}`);

		while (!player.lost && !player.hasSevenCards()) {
			const choice = await this.input.ask(
				`${player.name}, tirer une carte ? (o/n) `
			);

			this.logger.log(`${player.name} choisit : ${choice}`);

			if (choice !== "o") {
				console.log(`${player.name} s'arrête`);
				this.logger.log(`${player.name} s'arrête`);
				break;
			}

			const card = this.deck.draw();
			console.log(`${player.name} tire ${card}`);
			this.logger.log(`${player.name} tire la carte ${card}`);

			player.addCard(card);

			if (player.lost) {
				console.log(`${player.name} a perdu (carte en double)`);
				this.logger.log(`${player.name} PERD (carte en double)`);
			}
		}

		if (player.hasSevenCards()) {
			console.log(`${player.name} a gagné avec 7 cartes !`);
			this.logger.log(`${player.name} GAGNE (7 cartes)`);
		}
	}

	showScores() {
		console.log(" Scores finaux :");
		this.logger.log(" Scores finaux");

		for (const player of this.players) {
			const score = player.getScore();
			console.log(`${player.name} : ${score} points`);
			this.logger.log(`${player.name} : ${score} points`);
		}
	}
}
