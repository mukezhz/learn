import { ai } from './config';
import { interactiveChat } from './interactive';
import { menuSuggestionFlowMarkdown } from './menu_suggestion_flow';
import { addCalucator, finalizerFunction, getWeather } from './tool';

async function main() {
  // // basic
  // await (async () => {
  //   // make a generation request
  //   const { text, usage } = await ai.generate({
  //     system: "you are a flirty chatbot.",
  //     prompt: "Tell me a joke.",
  //   });
  //   console.log(text, usage);
  // })()

  // // flow
  // await (async () => {
  //   const text = await menuSuggestionFlowMarkdown('bistro');
  //   console.log(text);
  // })()

  // // template
  // await (async () => {
  //   const input = {
  //     location: "Nepal"
  //   }
  //   const menuPrompt = ai.prompt("menu");
  //   const { text, usage } = await menuPrompt(input);
  //   console.log(text, usage);
  // }
  // )()

  // // interactive chat
  // await (async () => {
  //   interactiveChat();
  // })()

  // tool
  await (async () => {
    const response = await ai.generate({
      system: `You are a chatbot that can provide the weather and do calculations. Finally always call the finalizer function. incase you are not able to generate the answer call the finalizer function by passing "unable to generate the answer"`,
      // prompt: 'What is the weather in Baltimore?',
      // prompt: 'What is the sum of 100 and 99?',
      prompt: 'Who is sushant babu?',
      tools: [getWeather, addCalucator, finalizerFunction],
    });
    console.log(response.text);
  })()
}

main();

