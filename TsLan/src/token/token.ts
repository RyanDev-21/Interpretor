export const TokenType = {
    ILLEGAL: "ILLEGAL",
    EOF: "EOF",
    IDENT: "IDENT",
    INT: "INT",
    FLOAT: "FLOAT",
    EQUAL: "==",
    NEQUAL: "!=",
    ASSIGN: "=",
    PLUS: "+",
    COMMA: ",",
    SEMICOLON: ";",
    LPAREN: "(",
    RPAREN: ")",
    LBRACE: "{",
    RBRACE: "}",
    MINUS: "-",
    SLASH: "/",
    LT: "<",
    GT: ">",
    LTEQUAL: "<=",
    GTEQUAL: ">=",
    BANG: "!",
    ASTERISK: "*",
    FUNCTION: "FUNCTION",
    LET: "LET",
    IF: "IF",
    ELSE: "ELSE",
    RETURN: "RETURN",
} as const;

export type TokenItem = (typeof TokenType)[keyof typeof TokenType];

const keywords = {
    "fn": TokenType.FUNCTION,
    "if": TokenType.IF,
    "else": TokenType.ELSE,
    "return": TokenType.RETURN,
    "let": TokenType.LET,
} as const;
type keywordsKey = keyof typeof keywords;
export type Token = {
    type: TokenItem,
    literal: string,
}

export function readIDENT(literal: string): TokenItem {
    let key = literal as keywordsKey
    switch (key) {
        case "let":
            return keywords.let
        case "fn":
            return keywords.fn
        case "return":
            return keywords.return
        case "if":
            return keywords.if
        case "else":
            return keywords.else
        default:
            return TokenType.IDENT
    }
}


export function identNumber(literal: string): TokenItem {
    let count = literal.match(/\./g)?.length || [].length;
    switch (count) {
        case 0:
            return TokenType.INT
        case 1:
            return TokenType.FLOAT
        default:
            return TokenType.ILLEGAL
    }
}




