package parser

import (
    "monke/ast"
    "monke/lexer"
    "monke/token"
)

type Parser struct {
    l *lexer.Lexer

    curToken token.Token
    peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    // Read two tokens, so curToken and peekToken are both set
    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) nextToken() {
    p.curToken = p.peekToken;
    p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for p.curToken.Literal != token.EOF {
        stmt := p.parseStatement()
        p.nextToken()

        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
    }

    return program
}

func (p *Parser) parseStatement() ast.Statement {
    switch (p.curToken.Type) {
    case token.LET:
        return p.parseLetStatement()
    default:
        return nil
    }
}

/*
func (p *Parser) parseLetStatement() ast.LetStatement {
    p.nextToken() // we advance the tokens because when this function is called we are at "let" keyword

    identifier := p.curToken.Literal

    


}
*/




