package lexer

import (
	"monke/token"
	"time"
)

type Lexer struct {
    input string
    position int // current position in input (points to current char)
    readPosition int // current reading position in input (after current char)
    ch byte // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{
        input: input,
    };

    l.readChar();
        
    return l;
}


// TODO: support for floats, negative ints, negative floats, infinite
func (l *Lexer) NextToken() token.Token {
    var tok token.Token;

    l.skipWhitespace();

    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            ch := l.ch;

            l.readChar();

            tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)};
        } else {
            tok = newToken(token.ASSIGN, l.ch);
        }
    case '+':
        tok = newToken(token.PLUS, l.ch);
    case '-':
        tok = newToken(token.MINUS, l.ch);
    case '!':
        if l.peekChar() == '=' {
            ch := l.ch;

            l.readChar();

            tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)};
        } else {
            tok = newToken(token.BANG, l.ch);
        }
    case '*':
        tok = newToken(token.ASTERISK, l.ch);
    case '/':
        tok = newToken(token.SLASH, l.ch);
    case '<':
        tok = newToken(token.LT, l.ch);
    case '>':
        tok = newToken(token.GT, l.ch);
    case '(':
        tok = newToken(token.LPAREN, l.ch);
    case ')':
        tok = newToken(token.RPAREN, l.ch);
    case '{':
        tok = newToken(token.LBRACE, l.ch);
    case '}':
        tok = newToken(token.RBRACE, l.ch);
    case ',':
        tok = newToken(token.COMMA, l.ch);
    case ';':
        tok = newToken(token.SEMICOLON, l.ch);
    case 0:
        tok.Type = token.EOF;
        tok.Literal = "";
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier();
            tok.Type = token.LookupIdent(tok.Literal);

            // early return because l.readIdentifier() moves our readPosition and position
            // past the character of the current identifier
            return tok;
        } else if isDigit(l.ch) {
            tok.Literal = l.readNumber();
            tok.Type = token.INT;

            // early return because l.readIdentifier() moves our readPosition and position
            // past the character of the current identifier
            return tok;
        } else {
            tok = newToken(token.ILLEGAL, l.ch);
        }
    }

    l.readChar();
    return tok;
}

// increments the readPosition and assigns the next readable char in l.ch
func (l *Lexer) readChar() {
    time.Sleep(10000)
    if l.readPosition >= len(l.input) {
        l.ch = 0;
    } else {
        l.ch = l.input[l.readPosition];
    }
    
    l.position = l.readPosition;
    l.readPosition++;
}

// reads a string from readPosition until a non-letter character
func (l *Lexer) readIdentifier() string {
    position := l.position;

    for isLetter(l.ch) {
        l.readChar();
    }

    return l.input[position:l.position];
}

// reads a string composed of digits from readPosition until a non-digit character
func (l *Lexer) readNumber() string {
    position := l.position;

    for isDigit(l.ch) {
        l.readChar();
    }

    return l.input[position:l.position];
}

// returns the next character that will be read next
func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0;
    }

    return l.input[l.readPosition];
}

// Returns a boolean specifying whether the given character is an integer.
// The function also looks at next characters to check whether this can be
// a legal int or not
// TODO: can I move this function out of the Lexer struct?
//func (l *Lexer) isInt() bool {
    // don't allow this type of expression: 013942
 //   if l.ch == '0' && !isEndToken(l.input[l.readPosition]) {
  //      return false;
   // }

    // don't allow this type of expressions: '-;' or '-.', etc.
    //if l.ch == '-' && !isDigit(l.input[l.readPosition]) {
     //   return false;
   // }

    //return l.ch >= '0' && l.ch <= '9';
//}

// skips any number of whitespace, tab, carriage return (Max OS pre-X), newline 
// characters between two string elements
func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
        l.readChar();
    }
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{
        Type: tokenType,
        Literal: string(ch),
    };
}

// controls what characters can be in identifiers
func isLetter(ch byte) bool {
    return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_';
}

// controls what elements are permitted in numbers
func isDigit(ch byte) bool {
    return ch >= '0' && ch <= '9';
}

