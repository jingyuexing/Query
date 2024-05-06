function isNumberic(ch: string) {
	return "0" <= ch && ch <= "9";
}
function isLetter(ch: string) {
	return ("a" <= ch && ch <= "z") || ("A" <= ch && ch <= "Z") || ch >= "\xff";
}

function isWhiteSpace(ch: string) {
	return "\n" === ch || " " === ch || ch === "\r" || ch === "\t";
}

export enum TokenKind {
	GreatThan = "GreatThan",
	LessThan = "LessThan",
	LessThanOrEqual = "LessThanOrEqual",
	GreatThanOrEqual = "GreatThanOrEqual",
	COLON = "Colon",
	Number = "Number",
	Text = "Text",
	Identifier = "Identifier",
	NotEqual = "NotEqual",
	Terminator = "Terminator",
	Comma = "Comma"
}
export interface Token {
	type: TokenKind;
	value: string;
}

function createToken(type: TokenKind, value: string) {
	return {
		type,
		value,
	};
}

// < =
// > =
// < >
// :
// repo:name

export function tokenizer(text: string) {
	let current = 0;
	const tokens: Token[] = [];
	let ch = text[current];
	function adavance() {
		current++;
		ch = text[current];
	}
	while (current < text.length) {
		ch = text[current];
		if (isWhiteSpace(ch)) {
			while(isWhiteSpace(ch) && current < text.length){
				adavance();
			}
			tokens.push(createToken(TokenKind.Terminator,";"))
			continue;
		}
		if (isNumberic(ch)) {
			let value = "";
			while ((isNumberic(ch) || ch === ".") && current <= text.length) {
				value += ch;
				adavance();
			}
			tokens.push(createToken(TokenKind.Number, value));
		}
		if (isLetter(ch)) {
			let value = "";
			while ((isLetter(ch) || ["-","$","_"].includes(ch)) && current <= text.length) {
				value += ch;
				adavance();
			}
			tokens.push(createToken(TokenKind.Identifier, value));
		}
		if (ch === ":") {
			tokens.push(createToken(TokenKind.COLON, ch));
			adavance()
			continue;
		}
		if(ch === ","){
			tokens.push(createToken(TokenKind.Comma,ch))
			adavance()
			continue;
		}
		if (ch === "<") {
			const slice = text.slice(current, current + 2);
			switch (slice) {
				case "<=":
					tokens.push(createToken(TokenKind.LessThanOrEqual, slice));
					current += 2;
					break;
				case "<>":
					tokens.push(createToken(TokenKind.NotEqual, slice));
					current += 2;
					break;
				default:
					tokens.push(createToken(TokenKind.LessThan, ch));
					adavance();
			}
		}
		if (ch === ">") {
			const slice = text.slice(current, current + 2);
			switch (slice) {
				case ">=":
					tokens.push(createToken(TokenKind.GreatThanOrEqual, slice));
					current +=2
					break;
				default:
					tokens.push(createToken(TokenKind.GreatThan, ch));
					adavance();
			}
		}
	}
	return tokens;
}