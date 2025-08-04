// Define a interface, deixando claro que 'console' é uma função
// que aceita um número variável de argumentos de qualquer tipo e retorna 'void'.
export interface ILibTS {
  console: (...data: any[]) => void;
}

export const libConsole = (...data: any[]): void => console.log(...data); 
// Nota: Adicionei '...' em 'console.log(...data)' para que os argumentos sejam
// espalhados, resultando em uma saída mais limpa no console.

// Porta de entrada para uso externo.
// O TypeScript agora sabe que a propriedade 'console' é uma função.
export const LibTS: ILibTS = {
  console: libConsole
};