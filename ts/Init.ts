export interface ILibTS {
  console: {
    libConsole: (...data: any[]) => void
  }
}

export const libConsole = (...data: any[]): void => console.log(...data) 

export const libsTS: ILibTS = {
  console: {
    libConsole
  }
}