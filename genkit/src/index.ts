import { ai } from './config';
import { interactiveChat } from './interactive';
import { menuSuggestionFlowMarkdown } from './menu_suggestion_flow';

async function main() {
  // basic
  await (async () => {
    // make a generation request
    const { text, usage } = await ai.generate({
      system: "you are a flirty chatbot.",
      prompt: "Tell me a joke.",
    });
    console.log(text, usage);
  })()

  // flow
  await (async () => {
    const text = await menuSuggestionFlowMarkdown('bistro');
    console.log(text);
  })()

  // template
  await (async () => {
    const input = {
      location: "Nepal"
    }
    const menuPrompt = ai.prompt("menu");
    const { text, usage } = await menuPrompt(input);
    console.log(text, usage);
  }
  )()

  // interactive chat
  await (async () => {
    interactiveChat();
  })()
}

main();

