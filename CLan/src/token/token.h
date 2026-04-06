#ifndef TOKEN_H
#define TOKEN_H

#include <stdint.h>

typedef enum {
    TOKEN_ASSIGN,
    TOKEN_PLUS,
    TOKEN_LPAREN,
    TOKEN_RPAREN,
    TOKEN_LBRACE,
    TOKEN_RBRACE,
    TOKEN_COMMA,
    TOKEN_SEMICOLON,
    TOKEN_EOF,
    TOKEN_ILLEGAL
} TokenType;

typedef struct {
    TokenType type;
    char literal;
} Token;

#endif
