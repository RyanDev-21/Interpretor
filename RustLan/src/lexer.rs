use crate::token::{Token, TokenType};

struct Lexer<'a> {
    input: &'a [u8],
    pos: usize,
    read_pos: usize,
    ch: u8,
}

impl<'a> Lexer<'a> {
    pub fn new_lexer(input: &'a [u8]) -> Self {
        let mut l = Lexer {
            input,
            pos: 0,
            read_pos: 0,
            ch: b'\0',
        };
        l.read_char();
        return l;
    }
    pub fn new_token(self: &Self, token_type: TokenType) -> Token<'a> {
        Token {
            token_type,
            literal: &self.input[self.pos..self.read_pos],
        }
    }

    pub fn next_token(self: &mut Self) -> Token<'a> {
        let t = match self.ch {
            b'=' => self.new_token(TokenType::ASSIGN),
            b'+' => self.new_token(TokenType::PLUS),
            b'(' => self.new_token(TokenType::LPAREN),
            b')' => self.new_token(TokenType::RPAREN),
            b'{' => self.new_token(TokenType::LBRACE),
            b'}' => self.new_token(TokenType::RBRACE),
            b',' => self.new_token(TokenType::COMMA),
            b';' => self.new_token(TokenType::SEMICOLON),
            0 => self.new_token(TokenType::EOF),
            _ => self.new_token(TokenType::ILLEGAL),
        };
        self.read_char();
        return t;
    }

    pub fn read_char(self: &mut Self) {
        if self.read_pos >= self.input.len() {
            self.ch = 0;
        } else {
            self.ch = self.input[self.read_pos];
        }
        self.pos = self.read_pos;
        self.read_pos += 1;
    }
}

#[cfg(test)]
mod test {
    use super::*; // Import Lexer, Token, TokenType from the outer file

    #[test]
    fn test_next_token() {
        let input = b"=+(){},;";

        // Define our expected output
        let tests = vec![
            (TokenType::ASSIGN, b"="),
            (TokenType::PLUS, b"+"),
            (TokenType::LPAREN, b"("),
            (TokenType::RPAREN, b")"),
            (TokenType::LBRACE, b"{"),
            (TokenType::RBRACE, b"}"),
            (TokenType::COMMA, b","),
            (TokenType::SEMICOLON, b";"),
            (TokenType::EOF, &[0u8]),
        ];

        let mut l = Lexer::new_lexer(input);

        for (expected_type, expected_literal) in tests {
            let tok = l.next_token();

            assert_eq!(tok.token_type, expected_type);
            assert_eq!(tok.literal, expected_literal);
        }
    }
}
