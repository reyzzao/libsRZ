// @file: Init.ts
import type { ILibTS } from "./Init.d.ts"; // A importação de IMessage foi removida, pois ela já está incluída na ILibTS
import { Log, Messages } from "./libs/LoaderLibs.ts";

// Porta de entrada
export const libTS: ILibTS = {
  Log,
  Messages
};
