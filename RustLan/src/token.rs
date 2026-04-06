#[derive(Debug, PartialEq)]
pub enum TokenType {
    ILLEGAL,
    EOF,
    // IDNET,
    // INT,
    ASSIGN,
    PLUS,
    COMMA,
    SEMICOLON,
    LPAREN,
    RPAREN,
    LBRACE,
    RBRACE,
    // FUNCTON,
    // LET,
}
#[derive(Debug, PartialEq)]
pub struct Token<'a> {
    pub token_type: TokenType,
    pub literal: &'a [u8],
}
