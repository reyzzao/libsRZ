
export interface LibTS {
  console: (...data: any[]) => void

  // test: string
  // fn: (d: string) => string
}

const testFN = (d: string) => d

// Porta de entrada
export const LibTS: LibTS = {
  console: (...data: any[]): void => console.log(...data), // console aqui evita erros


  // test: "foo1000",
  // fn: testFN,
}