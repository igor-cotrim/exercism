//https://exercism.org/tracks/typescript/exercises/resistor-color-trio

export function decodedResistorValue(colors: string[]): string {
  const duoColors = colors.slice(0, 2);
  const lastColor = colors[2];
  let value =
    (COLORS.indexOf(duoColors[0]) * 10 + COLORS.indexOf(duoColors[1])) *
    10 ** COLORS.indexOf(lastColor);

  const [divisor, prefix] = OHMS.find(([divisor]) => value >= divisor) ?? [
    1,
    "",
  ];

  console.log(`${value / divisor} ${prefix}ohms`);

  return `${value / divisor} ${prefix}ohms`;
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

const OHMS = [
  [1000000000, "giga"],
  [1000000, "mega"],
  [1000, "kilo"],
] as const;

decodedResistorValue(["orange", "orange", "black"]); // 33 ohms ok
decodedResistorValue(["blue", "grey", "brown"]); // 680 ohms ok
decodedResistorValue(["black", "black", "black"]); // 0 ohms ok
decodedResistorValue(["black", "grey", "black"]); // 8 ohms ok

decodedResistorValue(["red", "black", "red"]); // 2 kiloohms
decodedResistorValue(["green", "brown", "orange"]); // 51 kiloohms
decodedResistorValue(["yellow", "violet", "yellow"]); // 470 kiloohms
decodedResistorValue(["blue", "violet", "blue"]); // 67 megaohms
decodedResistorValue(["white", "white", "white"]); // 99 gigaohms
decodedResistorValue(["blue", "green", "yellow", "orange"]); // 650 kiloohms
