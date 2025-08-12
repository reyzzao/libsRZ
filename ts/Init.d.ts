// @file: ts/Init.d.ts
import { IMessage } from "./libs/LoaderLibs.ts";

export interface ILibTS {
  Log: (...data: any[]) => void;
  Messages: IMessage;
}

export declare const libTS: ILibTS;
