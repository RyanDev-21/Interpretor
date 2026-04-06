import { Lexer } from "./lexer"
import { Token, TokenType } from "../token/token";
describe("Test input", () => {
    it("should out the tokens", () => {
        const input: string = '=+(){},;';
        let lexer: Lexer = new Lexer(input);
        const test: Token[] = [
            { type: TokenType.ASSIGN, literal: "=" },
            { type: TokenType.PLUS, literal: "+" },
            { type: TokenType.LPAREN, literal: "(" },
            { type: TokenType.RPAREN, literal: ")" },
            { type: TokenType.LBRACE, literal: "{" },
            { type: TokenType.RBRACE, literal: "}" },
            { type: TokenType.COMMA, literal: "," },
            { type: TokenType.SEMICOLON, literal: ";" },
            { type: TokenType.EOF, literal: "" },
        ] as Token[];
        test.forEach(t => {
            let tok: Token = lexer.nextToken()
            expect(tok.type).toEqual(t.type)
            expect(tok.literal).toEqual(t.literal)
        });

    })
})
