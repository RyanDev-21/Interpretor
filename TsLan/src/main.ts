import { startRepl } from "./repl/repl"

const PROMPT = ">>";
console.log(`Welcom to the monkye langugae .${process.env.USERNAME || process.env.USER}`)
process.stdout.write(PROMPT)
process.stdin.on('data', (data) => {
    const input = data.toString().trim();

    if (input === 'exit') {
        process.exit(0);
    }

    startRepl(input)
    process.stdout.write(PROMPT);
})






