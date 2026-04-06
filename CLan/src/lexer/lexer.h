#ifndef Lexer_H
#define Lexer_H

#include "../token/token.h"
#include <stdlib.h>
typedef struct {
    char *input;
    size_t pos;
    size_t readPos;
    char ch;
} Lexer;
Lexer *NewLexer(char *input);
Token *nextToken(Lexer *l);
void freeLexer(Lexer *l);
#endif
