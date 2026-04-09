import { Lexer } from "../lexer/lexer";
import { TokenType } from "../token/token";



export function startRepl(input: string) {
    const l = new Lexer(input);
    for (let tok = l.nextToken(); tok.type != TokenType.EOF; tok = l.nextToken()) {
        console.log(tok)
    }
};

