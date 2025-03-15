import { genkit } from "genkit/beta";
import { googleAI, gemini20Flash } from "@genkit-ai/googleai";

import { createInterface } from "node:readline/promises";

const ai = genkit({
    plugins: [googleAI()],
    model: gemini20Flash,
});

export async function interactiveChat() {
    const chat = ai.chat();
    console.log("You're chatting with Gemini. Ctrl-C to quit.\n");
    const readline = createInterface(process.stdin, process.stdout);
    while (true) {
        const userInput = await readline.question("> ");
        const { text } = await chat.send(userInput);
        console.log(text);
    }
}