import { TokenItem, TokenType, Token } from "../token/token";

export class Lexer {
    private input: string
    private pos: number
    private readPos: number
    private char: string

    constructor(input: string) {
        this.input = input
        this.pos = 0
        this.readPos = 0
        this.char = input[this.readPos]
        this.readChar()
    }


    private readChar() {
        if (this.readPos >= this.input.length) {
            this.char = '\0'
        } else {
            this.char = this.input[this.readPos]
        }

        this.pos = this.readPos
        this.readPos++
    }


    public nextToken(): Token {
        var tok: Token
        switch (this.char) {
            case '=':
                tok = this.newToken(TokenType.ASSIGN, this.char)
                break
            case ';':
                tok = this.newToken(TokenType.SEMICOLON, this.char)
                break
            case '(':
                tok = this.newToken(TokenType.LPAREN, this.char)
                break;
            case ')':
                tok = this.newToken(TokenType.RPAREN, this.char)
                break
            case ',':
                tok = this.newToken(TokenType.COMMA, this.char)
                break
            case '+':
                tok = this.newToken(TokenType.PLUS, this.char)
                break
            case '{':
                tok = this.newToken(TokenType.LBRACE, this.char)
                break
            case '}':
                tok = this.newToken(TokenType.RBRACE, this.char)
                break
            case '\0':
                tok = this.newToken(TokenType.EOF, "")
                break
            default:
                tok = this.newToken(TokenType.ILLEGAL, "")
                break
        }
        this.readChar()
        return tok

    }

    private newToken(TokenType: TokenItem, char: string): Token {
        return {
            type: TokenType,
            literal: char
        } as Token
    }

}
