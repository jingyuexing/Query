package query

import (
	"strings"
	"unicode"

)

type TokenKind string

const (
	GreatThan        TokenKind = "GreatThan"
	LessThan         TokenKind = "LessThan"
	LessThanOrEqual  TokenKind = "LessThanOrEqual"
	GreatThanOrEqual TokenKind = "GreatThanOrEqual"
	COLON            TokenKind = "Colon"
	TNumber          TokenKind = "Number"
	TText            TokenKind = "Text"
	Identifier       TokenKind = "Identifier"
	NotEqual         TokenKind = "NotEqual"
	Comma            TokenKind = "Comma"
	Quote            TokenKind = "Quote"
	Terminator       TokenKind = "Terminator"
)

type Token struct {
	Type  TokenKind
	Value string
}

func isValidCharacter(ch rune) bool {
	return unicode.IsLetter(ch) || ch > 0xff
}

func createToken(kind TokenKind, value string) Token {
	return Token{Type: kind, Value: value}
}

func Tokenizer(text string) []Token {
	text_ := []rune(text)
	length := len(text_)
	current := 0
	tokens := make([]Token, 0)
	var ch rune
	ch = rune(text_[current])
	advance := func() {
		current++
		if current >= length {
			ch = text_[length -1]
		}else{
			ch = text_[current]
		}	
	}

	if length > 0 {
		ch = rune(text_[current])
	}
	for current < length {
		if unicode.IsNumber(rune(ch)) {
			value := ""
			dot := 0
			for (unicode.IsNumber(rune(ch)) || ch == '.') && current < length {
				value += string(ch)
				if ch == '.' {
					dot++
				}
				advance()
			}
			if dot <= 1 {
				tokens = append(tokens, createToken(TNumber, value))
			} else {
				tokens = append(tokens, createToken(TText, value))

			}
		}

		if isValidCharacter(rune(ch)) {
			value := ""
			for (isValidCharacter(rune(ch)) || unicode.IsNumber(rune(ch)) || strings.ContainsRune("-$_", rune(ch))) && current < length {
				value += string(ch)
				advance()
			}
			tokens = append(tokens, createToken(Identifier, value))
		}
		if ch == ' ' {
			for ch == ' ' {
				advance()
			}
			tokens = append(tokens, createToken(Terminator, ";"))
		}
		if ch == ':' {
			tokens = append(tokens, createToken(COLON, string(ch)))
			advance()
			continue
		}

		if ch == ',' {
			tokens = append(tokens, createToken(Comma, string(ch)))
			advance()
			continue
		}

		if ch == '<' {
			slice := text[current:min(current+2, length)]
			if slice == "<=" {
				tokens = append(tokens, createToken(LessThanOrEqual, slice))
				current += 2
			} else if slice == "<>" {
				tokens = append(tokens, createToken(NotEqual, slice))
				current += 2
			} else {
				tokens = append(tokens, createToken(LessThan, string(ch)))
				advance()
			}
		}

		if ch == '>' {
			slice := text[current:min(current+2, length)]
			if slice == ">=" {
				tokens = append(tokens, createToken(GreatThanOrEqual, slice))
				current += 2
			} else {
				tokens = append(tokens, createToken(GreatThan, string(ch)))
				advance()
			}
		}
		if ch == '"' || ch == '\'' {
			end := ch
			val := ""
			advance()
			for ch != end && current < length {
				val = val + string(ch)
			}
			tokens = append(tokens, createToken(Quote, val))
		}
	}
	return tokens
}
