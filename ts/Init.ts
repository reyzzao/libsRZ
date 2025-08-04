
export interface LibTS {
  test: string,
  fn: (d: string) => string,
  console: (...data: any[]) => void
}

const testFN = (d: string) => d

// Porta de entrada
export const LibTS: LibTS = {
  test: "foo1000",
  fn: testFN,
  console: (...data: any[]): void => console.log(...data), // console aqui evita erros
}