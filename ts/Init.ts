import { libConsole } from "./console/Lib"

export interface ILibTS {
  console: void
}

// Porta de entrada para uso externo.
export const LibTS: ILibTS = {
  console: libConsole()
}

// console.log(LibTS)