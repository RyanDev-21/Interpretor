#include "lexer/lexer.h"
#include "token/token.h"
#include <assert.h>
#include <stdio.h>
typedef struct {
    TokenType expectedType;
    char expectdliteral;
} ExpectTokenType;

int main() {
    Lexer *l = NewLexer("+=(){},;");
    ExpectTokenType tests[] = {{TOKEN_PLUS, '+'},   {TOKEN_ASSIGN, '='},    {TOKEN_LPAREN, '('},
                               {TOKEN_RPAREN, ')'}, {TOKEN_LBRACE, '{'},    {TOKEN_RBRACE, '}'},
                               {TOKEN_COMMA, ','},  {TOKEN_SEMICOLON, ';'}, {TOKEN_EOF, '\0'}};

    int length_test = sizeof(tests) / sizeof(ExpectTokenType);
    for (int i = 0; i < length_test; i++) {
        Token *t = nextToken(l);
        assert(t->type == tests[i].expectedType);
        assert(t->literal == tests[i].expectdliteral);
        free(t);
    }
    printf("All tests passed\n");
    freeLexer(l);
    return 0;
}
