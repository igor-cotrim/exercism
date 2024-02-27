//https://exercism.org/tracks/typescript/exercises/resistor-color-duo

export function decodedValue(colors: string[]): number {
  let value = 0;

  const duoColors = colors.slice(0, 2);

  for (let i = 0; i < duoColors.length; i++) {
    value += COLORS.indexOf(colors[i]) * Math.pow(10, duoColors.length - i - 1);
  }

  return value;
}

export const COLORS = [
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white",
];
