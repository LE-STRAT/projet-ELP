export default class Player {
	constructor(name) {
		this.name = name;
		this.cards = [];
		this.lost = false;
	}

	addCard(card) {
		if (this.cards.includes(card)) {
			this.lost = true;
		} else {
			this.cards.push(card);
		}
	}

	hasSevenCards() {
		return this.cards.length === 7;
	}

	getScore() {
		if (this.lost) return 0;

		return this.cards.reduce((sum, card) => sum + card, 0);
	}
}
