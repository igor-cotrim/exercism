// https://exercism.org/tracks/typescript/exercises/space-age

export function age(planet: string, seconds: number): number {
  let year = 0;

  switch (planet.toLowerCase()) {
    case "mercury":
      year = 0.2408467;
      break;
    case "venus":
      year = 0.61519726;
      break;
    case "earth":
      year = 1.0;
      break;
    case "mars":
      year = 1.8808158;
      break;
    case "jupiter":
      year = 11.862615;
      break;
    case "saturn":
      year = 29.447498;
      break;
    case "uranus":
      year = 84.016846;
      break;
    case "neptune":
      year = 164.79132;
      break;
    default:
      throw new Error("Invalid planet.");
  }

  return Number((seconds / (year * 31557600)).toFixed(2));
}
