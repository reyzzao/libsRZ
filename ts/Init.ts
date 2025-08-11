// @file: Init.ts
import type { ILibTS } from "./Init.d.ts"; // A importação de IMessage foi removida, pois ela já está incluída na ILibTS
import { Log } from "./libs/console/Index.ts";
import { message } from "./libs/messages/Mod.ts";

// Porta de entrada
export const libTS: ILibTS = {
  Log,
  message,
};
