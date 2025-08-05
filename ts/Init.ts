// @file: Init.ts
import type { ILibTS } from "./Init.d.ts" // A importação de IMessage foi removida, pois ela já está incluída na ILibTS
import { message } from "./libs/messages/Mod.ts"


// Porta de entrada
export const libTS: ILibTS = {
  console: (...data: any[]): void => console.log(...data), // console aqui evita erros
  message,
}