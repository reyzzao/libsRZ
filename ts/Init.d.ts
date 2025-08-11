// @file: Init.d.ts
import { IMessage } from "./libs/messages/Mod.ts";

export interface ILibTS {
  Log: (...data: any[]) => void
  message:IMessage
}

export declare const libTS: ILibTS;