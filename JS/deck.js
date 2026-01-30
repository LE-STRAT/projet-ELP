export default class Deck {
	constructor() {
		this.cards = [];
		this.init();
		this.shuffle();
	}

	init() {
		for (let value = 1; value <= 12; value++) {
			for (let i = 0; i < value; i++) {
				this.cards.push(value);
			}
		}
	}

	shuffle() {
		this.cards.sort(() => Math.random() - 0.5);
	}

	draw() {
		if (this.cards.length === 0) {
			return null;
		}
		return this.cards.pop();
	}
}
