import { z } from "genkit";
import { ai } from "./config";
import { gemini20Flash } from "@genkit-ai/googleai";

const MenuItemSchema = z.object({
    dishname: z.string(),
    description: z.string(),
});


export const menuSuggestionFlowMarkdown = ai.defineFlow(
    {
        name: 'menuSuggestionFlow',
        inputSchema: z.string(),
        outputSchema: z.string(),
    },
    async (restaurantTheme) => {
        const { output } = await ai.generate({
            model: gemini20Flash,
            prompt: `Invent a menu item for a ${restaurantTheme} themed restaurant.`,
            output: { schema: MenuItemSchema },
        });
        if (output == null) {
            throw new Error("Response doesn't satisfy schema.");
        }
        return `**${output.dishname}**:\n${output.description}`;
    }
);