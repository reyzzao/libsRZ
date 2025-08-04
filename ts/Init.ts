// import { libConsole } from './console/Lib.ts'


// export interface ILibTS {
//   console: {
//     libConsole: (...data: any[]) => void
//   }
// }

export const libConsole = (...data: any[]): void => console.log(...data) 

// Porta de entrada
export const LibTS = {
  test: "foo1000",
  console: libConsole
}