export default class Player {
	constructor(name) {
		this.name = name;
		this.cards = [];
		this.lost = flase;
	}

	addCard(card) {
		if (this.cards.includes(cards)) {
			this.lost = true;
		} else {
			this.cards.push(card);
		}
	}

	hasSevenCards() {
		return this.cards.length === 7;
	}
}
