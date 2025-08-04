import { libConsole } from './console/Lib.ts'


export interface ILibTS {
  console: {
    libConsole: (...data: any[]) => void
  }
}


// Porta de entrada
export const libsTS: ILibTS = {
  console: {
    libConsole
  }
}