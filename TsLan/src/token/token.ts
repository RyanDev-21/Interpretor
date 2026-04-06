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
    FUNCTION: "FUNCTION", LET: "LET",
} as const;

export type TokenItem = (typeof TokenType)[keyof typeof TokenType];

export interface Token {
    type: TokenItem;
    literal: string;
}

