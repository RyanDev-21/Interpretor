import { TokenItem, TokenType, Token, readIDENT, identNumber } from "../token/token";



const a = 'a'.charCodeAt(0)
const A = 'A'.charCodeAt(0)
const z = 'z'.charCodeAt(0)
const Z = 'Z'.charCodeAt(0)
const _ = '_'.charCodeAt(0)
const _0 = '0'.charCodeAt(0)
const _9 = '9'.charCodeAt(0)
const dot = '.'.charCodeAt(0)
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

        this.skipWhiteSpace();
        switch (this.char) {
            case '=':
                if (this.peekChar() == "=") {
                    tok = this.newToken(TokenType.EQUAL, this.input[this.pos] + this.input[this.readPos])
                } else {
                    tok = this.newToken(TokenType.ASSIGN, this.char)
                }
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
            case '*':
                tok = this.newToken(TokenType.ASTERISK, this.char)
                break
            case '-':
                tok = this.newToken(TokenType.MINUS, this.char)
                break
            case '<':
                if (this.peekChar() == "=") {
                    tok = this.newToken(TokenType.LTEQUAL, this.input[this.pos] + this.input[this.readPos])
                } else {
                    tok = this.newToken(TokenType.LT, this.char)
                }
                break
            case '>':
                if (this.peekChar() == "=") {
                    tok = this.newToken(TokenType.GTEQUAL, this.input[this.pos] + this.input[this.readPos])

                } else {
                    tok = this.newToken(TokenType.GT, this.char)
                }
                break
            case '!':
                if (this.peekChar() == "=") {
                    tok = this.newToken(TokenType.NEQUAL, this.input[this.pos] + this.input[this.readPos])
                } else {
                    tok = this.newToken(TokenType.BANG, this.char)
                }
                break
            case '/':
                tok = this.newToken(TokenType.SLASH, this.char)
                break
            case '\0':
                tok = this.newToken(TokenType.EOF, "")
                break
            default:
                if (this.isLetter(this.char.charCodeAt(0))) {
                    const literal = this.readIdentifier()
                    const type = readIDENT(literal)
                    tok = this.newToken(type, literal)
                    return tok
                } else if (this.isDigit(this.char.charCodeAt(0))) {
                    const literal = this.readNumber()
                    const type = identNumber(literal)
                    tok = this.newToken(type, literal)
                    return tok
                }
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
    private readIdentifier(): string {
        let position = this.pos;
        while (this.isLetter(this.char.charCodeAt(0))) {
            this.readChar()
        }

        return this.input.slice(position, this.pos)
    }

    private readNumber(): string {
        let position = this.pos;
        while (this.isDigit(this.char.charCodeAt(0))) {
            this.readChar()
        }

        return this.input.slice(position, this.pos)
    }


    private peekChar(): string {
        if (this.readPos >= this.input.length) {
            return '\0';
        }
        return this.input[this.readPos]
    }


    private isLetter(char: number): boolean {
        return char >= A && char <= Z || char >= a && char <= z || char === _
    }

    private skipWhiteSpace() {
        while (this.char == ' ' || this.char == '\t' || this.char == '\n' || this.char == '\r') {
            this.readChar()
        }
    }

    private isDigit(char: number): boolean {
        return char >= _0 && char <= _9 || char === dot
    }

}
