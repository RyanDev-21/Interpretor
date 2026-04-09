import { Token, TokenType } from "../token/token";
import { Lexer } from "./lexer"
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
    it("testing the tokenizer phase 2 (Identifiers and Operators)", () => {
        const input = `let five = 5.5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;
`;

        const tests: Token[] = [
            { type: TokenType.LET, literal: "let" },
            { type: TokenType.IDENT, literal: "five" },
            { type: TokenType.ASSIGN, literal: "=" },
            { type: TokenType.FLOAT, literal: "5.5" }, // Note: You'll need an isDigit check for this!
            { type: TokenType.SEMICOLON, literal: ";" },
            { type: TokenType.LET, literal: "let" },
            { type: TokenType.IDENT, literal: "ten" },
            { type: TokenType.ASSIGN, literal: "=" },
            { type: TokenType.INT, literal: "10" },
            { type: TokenType.SEMICOLON, literal: ";" },
            { type: TokenType.LET, literal: "let" },
            { type: TokenType.IDENT, literal: "add" },
            { type: TokenType.ASSIGN, literal: "=" },
            { type: TokenType.FUNCTION, literal: "fn" },
            { type: TokenType.LPAREN, literal: "(" },
            { type: TokenType.IDENT, literal: "x" },
            { type: TokenType.COMMA, literal: "," },
            { type: TokenType.IDENT, literal: "y" },
            { type: TokenType.RPAREN, literal: ")" },
            { type: TokenType.LBRACE, literal: "{" },
            { type: TokenType.IDENT, literal: "x" },
            { type: TokenType.PLUS, literal: "+" },
            { type: TokenType.IDENT, literal: "y" },
            { type: TokenType.SEMICOLON, literal: ";" },
            { type: TokenType.RBRACE, literal: "}" },
            { type: TokenType.SEMICOLON, literal: ";" },
            { type: TokenType.LET, literal: "let" },
            { type: TokenType.IDENT, literal: "result" },
            { type: TokenType.ASSIGN, literal: "=" },
            { type: TokenType.IDENT, literal: "add" },
            { type: TokenType.LPAREN, literal: "(" },
            { type: TokenType.IDENT, literal: "five" },
            { type: TokenType.COMMA, literal: "," },
            { type: TokenType.IDENT, literal: "ten" },
            { type: TokenType.RPAREN, literal: ")" },
            { type: TokenType.SEMICOLON, literal: ";" },
            { type: TokenType.BANG, literal: "!" },
            { type: TokenType.MINUS, literal: "-" },
            { type: TokenType.SLASH, literal: "/" },
            { type: TokenType.ASTERISK, literal: "*" },
            { type: TokenType.INT, literal: "5" },
            { type: TokenType.SEMICOLON, literal: ";" },
            { type: TokenType.INT, literal: "5" },
            { type: TokenType.LT, literal: "<" },
            { type: TokenType.INT, literal: "10" },
            { type: TokenType.RT, literal: ">" },
            { type: TokenType.INT, literal: "5" },
            { type: TokenType.SEMICOLON, literal: ";" },
            { type: TokenType.EOF, literal: "" },
        ];

        const lexer = new Lexer(input);

        tests.forEach((expectedToken, i) => {
            const tok = lexer.nextToken();

            // Debugging tip: if it fails, show which index failed
            try {
                expect(tok.type).toEqual(expectedToken.type);
                expect(tok.literal).toEqual(expectedToken.literal);
            } catch (e) {
                throw new Error(`Test failed at index ${i}: expected ${expectedToken.type}, got ${tok.type}`);
            }
        });
    });
})
