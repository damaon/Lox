package ast

import "fmt"

func Scan(source string) ([]Token, error) {
    runes := []rune(source)

	start := 0
	current := 0
	line := 1

	tokens := make([]Token, 0)

	isEnd := func() bool {
		return current >= len(runes)
	}

	peek := func() rune {
		if isEnd(){
			return '\x00'
		}

		return runes[current]
	}

	peekNext := func() rune {
		if current + 1 >= len(runes){
			return '\x00'
		}

		return runes[current]
	}

	advance := func() rune {
		current++;
		return runes[current-1]
	}

	isNext := func(r rune) bool {
		if isEnd() || runes[current] != r {
			return false
		}

		current++;

		return true;
	}

	isDigit := func(r rune) bool {
		if r >= '0' && r <= '9' {
			return true
		}

		return false
	}

	isLetter := func(r rune) bool {
		if(r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'){
			return true
		}

		return false
	}

	addToken := func(typ TokenType){
		tokens = append(tokens, Token { typ, string(runes[start:current]), "", line })
	}

	scanToken := func() error {
		r := advance()
		switch r {
		case ' ':
		case '\t':
		case '\r':
			{
				break // ignore white space
			}

		case '\n':
			{
				line++
				break
			}

		// Single-character lexeme: '(', ')', '[', ']', '.', '-', '+', ';'
		case '(':
			{
				addToken(LEFT_PAREN)
				break
			}

		case ')':
			{
				addToken(RIGHT_PAREN)
				break
			}

		case '{':
			{
				addToken(LEFT_BRACE)
				break
			}

		case '}':
			{
				addToken(RIGHT_BRACE)
				break
			}

		case '.':
			{
				addToken(DOT)
				break
			}

		case '-':
			{
				addToken(MINUS)
				break
			}

		case '+':
			{
				addToken(PLUS)
				break
			}

		case '*':
			{
				addToken(STAR)
				break
			}

		case ',':
			{
				addToken(COMMA)
				break
			}

		case ';':
			{
				addToken(SEMICOLON)
				break
			}

		// Multi-character lexeme (potentially): '/', '!', '=', '<', '>', '!=', '==', '<=', '>=', '//'
		case '!':
			{
				if isNext('=') {
					addToken(BANG_EQUAL)
				} else {
					addToken(BANG)
				}

				break
			}

		case '=':
			{
				if isNext('=') {
					addToken(EQUAL_EQUAL)
				} else {
					addToken(EQUAL)
				}

				break
			}

		case '>':
			{
				if isNext('=') {
					addToken(GREATER_EQUAL)
				} else {
					addToken(GREATER)
				}

				break
			}

		case '<':
			{
				if isNext('=') {
					addToken(LESS_EQUAL)
				} else {
					addToken(LESS)
				}
			}

		case '/':
			{
				if isNext('/') {
					for peek() != '\n' && !isEnd() {
						advance()
					}
				} else {
					addToken(SLASH)
				}
			}

		case '"':
			{
				for peek() != '"' && !isEnd() {
					if peek() == '\n' {
						line++
					}

					advance()
				}

				// unterminated string
				if isEnd() {
					return fmt.Errorf("error at line %d: unterminated string", line)
				}

				advance()

				lexeme := string(runes[start:current])
				literal := string(runes[start+1 : current-1]) // remove double quotes

				tokens = append(tokens, Token{STRING, lexeme, literal, line})
			}

		default:
			{
				if isDigit(r) {
					for isDigit(peek()) {
						advance()
					}

					if peek() == '.' && isDigit(peekNext()) {
						advance()

						for isDigit(peek()) {
							advance()
						}
					}

					number := string(runes[start:current])
					tokens = append(tokens, Token{NUMBER, number, number, line})
				} else if isLetter(r) {
					for isLetter(peek()) || isDigit(peek()) {
						advance()
					}

					if t, ok := keywords[string(runes[start:current])]; ok {
						tokens = append(tokens, Token{t, string(runes[start:current]), "", line})
					} else {
						tokens = append(tokens, Token{IDENTIFIER, string(runes[start:current]), "", line})
					}
				} else {
					return fmt.Errorf("unknown character '%v' at line %d", string(r), line)
				}
			}
		}

		return nil
	}

	for !isEnd() {
		start = current
		if err := scanToken(); err != nil {
			return nil, err
		}
	}

	// cannot use addToken because lexeme will get the last character
	tokens = append(tokens, Token{EOF, "", "", line})

	return tokens, nil
}
