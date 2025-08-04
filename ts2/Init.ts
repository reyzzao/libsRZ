import { libConsole } from "./console/Lib"

interface ILibsList {
  console: void
}

const libsList: ILibsList = {
  console: libConsole()
}

function libTs(): ILibsList{
  return libsList
}

console.log(libTs())