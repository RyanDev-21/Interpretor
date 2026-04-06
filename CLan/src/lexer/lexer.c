#include "./lexer.h"
#include <stdlib.h>
#include <string.h>
static void read_char(Lexer *l);
static Token *newToken(TokenType type, char lit);

Lexer *NewLexer(char *input) {
    Lexer *l = malloc(sizeof(Lexer));
    l->input = input;
    l->pos = 0;
    l->readPos = 0;
    l->ch = '\0';
    read_char(l);
    return l;
};
static void read_char(Lexer *l) {
    if (l->readPos >= strlen(l->input)) {
        l->ch = '\0';
    } else {
        l->ch = l->input[l->readPos];
    }
    l->pos = l->readPos;
    l->readPos++;
};

Token *nextToken(Lexer *l) {
    Token *tok;
    switch (l->ch) {
    case '=':
        tok = newToken(TOKEN_ASSIGN, l->ch);
        break;
    case '+':
        tok = newToken(TOKEN_PLUS, l->ch);
        break;
    case '(':
        tok = newToken(TOKEN_LPAREN, l->ch);
        break;
    case ')':
        tok = newToken(TOKEN_RPAREN, l->ch);
        break;
    case '{':
        tok = newToken(TOKEN_LBRACE, l->ch);
        break;
    case '}':
        tok = newToken(TOKEN_RBRACE, l->ch);
        break;
    case ',':
        tok = newToken(TOKEN_COMMA, l->ch);
        break;
    case ';':
        tok = newToken(TOKEN_SEMICOLON, l->ch);
        break;
    case '\0':
        tok = newToken(TOKEN_EOF, '\0');
        break;
    default:
        tok = newToken(TOKEN_ILLEGAL, l->ch);
        break;
    }
    read_char(l);
    return tok;
};

static Token *newToken(TokenType type, char lit) {
    Token *t = malloc(sizeof(Token));
    t->type = type;
    t->literal = lit;
    return t;
}
void freeLexer(Lexer *l) { free(l); }
