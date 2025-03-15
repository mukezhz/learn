import { gemini20Flash, googleAI } from "@genkit-ai/googleai";
import { genkit } from "genkit";

export const ai = genkit({
    plugins: [googleAI()],
    model: gemini20Flash, // set default model
});