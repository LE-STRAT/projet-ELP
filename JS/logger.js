import fs from "fs";
import path from "path";

export default class Logger {
	constructor() {
		const date = new Date().toISOString().replace(/:/g, "-");
		this.filePath = path.join("logs", `game_${date}.txt`);

		if (!fs.existsSync("logs")) {
			fs.mkdirSync("logs");
		}

		fs.writeFileSync(this.filePath, "=== Nouvelle partie Flip 7 ===\n");
	}

	log(message) {
		const timestamp = new Date().toLocaleString();
		fs.appendFileSync(
			this.filePath,
		`[${timestamp}] ${message}\n`
		);
	}
}
