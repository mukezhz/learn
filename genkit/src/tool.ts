import { z } from 'genkit';
import { ai } from './config';

export const getWeather = ai.defineTool(
  {
    name: 'getWeather',
    description: 'Gets the current weather in a given location',
    inputSchema: z.object({
      location: z.string().describe('The location to get the current weather for')
    }),
    outputSchema: z.string(),
  },
  async (input) => {
    return `The current weather in ${input.location} is 63°F and sunny.`;
  }
);

export const addCalucator = ai.defineTool(
  {
    name: 'addCalculator',
    description: 'Performs add calculation',
    inputSchema: z.object({
      a: z.string().describe('The first operator'),
      b: z.string().describe('The second operator')
    }),
    outputSchema: z.number(),
  },
  async ({ a, b }) => {
    return Number(a) + Number(b);
  }
);

export const finalizerFunction = ai.defineTool(
  {
    name: 'finalizerFunction',
    description: 'Send the response in polite and formal way to the user. Apologies for any mistakes by explaining question.',
    inputSchema: z.object({
      a: z.string().describe('the final output of all tools'),
    }),
    outputSchema: z.string(),
  },
  async ({ a }) => {
    console.log('Finalizer function called:::', a);
    return `The final output is ${a}`;
  }
);

export const subtractCalculator = ai.defineTool(
  {
    name: 'subtractCalculator',
    description: 'The function which will subtract two numbers',
    inputSchema: z.object({
      a: z.number().describe('the first operand'),
      b: z.number().describe('the second operand'),
    }),
    outputSchema: z.string(),
  },
  async ({ a, b }) => {
    return `Result of substraction is ${a - b}`;
  }
);