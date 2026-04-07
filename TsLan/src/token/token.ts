export const TokenType = {
    ILLEGAL: "ILLEGAL",
    EOF: "EOF",
    IDENT: "IDENT",
    INT: "INT",
    ASSIGN: "=",
    PLUS: "+",
    COMMA: ",",
    SEMICOLON: ";",
    LPAREN: "(",
    RPAREN: ")",
    LBRACE: "{",
    RBRACE: "}",
    FUNCTION: "FUNCTION",
    LET: "LET",

} as const;

export type TokenItem = (typeof TokenType)[keyof typeof TokenType];

const keywords = {
    "fn": TokenType.FUNCTION,
    "let": TokenType.LET,
}
export class Tokenizer {
    private token_type: TokenItem;
    private literal: string;


    constructor(token_type: TokenItem, literal: string) {
        this.token_type = token_type;
        this.literal = literal;
    }

    public readIDENT(literal: string): TokenItem {
        switch (literal) {
            case keywords.let:
                return keywords.let
            case keywords.fn:
                return keywords.fn
            default:
                return TokenType.IDENT

        }
    }


}

