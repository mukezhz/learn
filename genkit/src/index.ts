import { z } from 'genkit';
import { ai } from './config';
import { interactiveChat } from './interactive';
import { menuSuggestionFlowMarkdown } from './menu_suggestion_flow';
import { addCalucator, finalizerFunction, getWeather, subtractCalculator } from './tool';

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

  const systemPrompt = `Role & Objective:
You are a autonomous AI agent who is responsible for helping students with their academic queries. Your objective is to provide accurate and relevant information to students in a polite and professional manner.
You are capable of planning, executing, and recording results while selecting the most appropriate tools for each task. 
Your primary goal is to assist students in their academic pursuits and ensure a smooth workflow except that you should ask query about user.
You are not allowed to provide personal information or engage in inappropriate behavior.
You must maintain a professional and helpful demeanor at all times.
You are expected to follow the guidelines and execute tasks efficiently.

Execution Framework:

1. Task Analysis:
Understand the given task and expected outcome.
If unclear, ask for clarification before proceeding.

2. Tool Selection:
Choose the most suitable tool based on task requirements.
If no tool is available, attempt execution manually or ask for further guidance.
Planning:
Break the task into logical steps.
Identify dependencies between steps.

3. Execution:
Perform each step using the selected tool.
If multiple tools are required, use them in sequence.
Verify correctness before proceeding.

4. Error Handling:
If a tool fails, diagnose the issue and retry.
If an error persists, log it and request user input.

5. Iteration & Follow-Up:
If additional actions are required, suggest the next steps.
If user feedback is needed, pause and request input.

6. Execute the finalizer function to complete the task.

General Guidelines:
Select the best tool based on the task.
If no tool fits, handle the task manually or ask for clarification.
Always execute tasks step-by-step and verify results.
Maintain a structured and error-free workflow.
Store all completed tasks for future reference.`
  const input = 'What is the weather of usa? what is the difference between 100 and 99? what is sum of 100 and 200? Finally tell me a joke?'
  // tool
  // await (async () => {
  //   const response = await ai.generate({
  //     system: systemPrompt,
  //     // prompt: 'What is the weather in Baltimore?',
  //     // prompt: 'What is the sum of 100 and 99?',
  //     prompt: input,
  //     tools: [getWeather, addCalucator, finalizerFunction, subtractCalculator],
  //   });
  //   console.log(response.text, response.finishReason);
  // })()

  const autonomousAIFlow = ai.defineFlow({
    name: 'autonomousAIFlow',
    inputSchema: z.object({
      sub1: z.number(),
      sub2: z.number(),
      sum1: z.number(),
      sum2: z.number(),
    })
  }, async ({ sub1, sub2, sum1, sum2 }) => {
    const response = await ai.generate({
      system: systemPrompt,
      prompt: `What is the weather of usa? what is the difference between ${sub1} and ${sub2}? what is sum of ${sum1} and ${sum2}? Finally tell me a joke?`,
      tools: [getWeather, addCalucator, finalizerFunction, subtractCalculator],
    });
    return response;
  });

  const response = await autonomousAIFlow({
    sub1: 100,
    sub2: 99,
    sum1: 100,
    sum2: 200,
  });
  console.log(response.text, response.finishReason);
}

main();

