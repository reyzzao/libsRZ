// import { libConsole } from './console/Lib.ts'


export interface LibTS {
  test: string,
  fn: (d: string) => string
}

// export const libConsole = (...data: any[]): void => console.log(...data) 
const testFN = (d: string) => d

// Porta de entrada
export const LibTS: LibTS = {
  test: "foo1000",
  fn: testFN
}