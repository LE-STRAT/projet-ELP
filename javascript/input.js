import readline from "readline";

export default class Input {
	constructor() {
		this.rl = readline.createInterface({
			input: process.stdin,
			output: process.stdout
		});
	}

	ask(question) {
		return new Promise(resolve => {
			this.rl.question(question, answer => {
				resolve(answer.trim());
			});
		});
	}

	close() {
		this.rl.close();
	}
}
