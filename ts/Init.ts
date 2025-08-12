// @file: ts/Init.ts
import type { ILibTS } from "./Init.d.ts"; // A importação de IMessage foi removida, pois ela já está incluída na ILibTS
import { Log, Messages } from "./libs/LoaderLibs.ts";
import { loaderModulesIndex } from "./libs/loader_index/LoaderModulesIndexFunction.ts";

// Porta de entrada
export const libTS: ILibTS = {
  Log,
  Messages,
  loaderModulesIndex,
};
