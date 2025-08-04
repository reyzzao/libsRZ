import { libConsole } from "./console/Lib"

export interface ILibTS {
  console: void
}

const libsList: ILibTS = {
  console: libConsole()
}

export function LibTS(): ILibTS{
  return libsList
}

console.log(LibTS())