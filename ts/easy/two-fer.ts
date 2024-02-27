//https://exercism.org/tracks/typescript/exercises/two-fer

/**
 * This stub is provided to make it straightforward to get started.
 */

export function twoFer(name?: string): string {
  if (name == "" || name == null) {
    return "One for you, one for me.";
  }

  return `One for ${name}, one for me.`;
}
