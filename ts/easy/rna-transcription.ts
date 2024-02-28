// https://exercism.org/tracks/typescript/exercises/rna-transcription

export function toRna(dna: string): string {
  let rna = "";

  for (let i = 0; i < dna.length; i++) {
    switch (dna[i]) {
      case "G":
        rna += "C";
        break;
      case "C":
        rna += "G";
        break;
      case "T":
        rna += "A";
        break;
      case "A":
        rna += "U";
        break;
      default:
        throw new Error("Invalid input DNA.");
    }
  }

  return rna;
}
