import { parser, type RootNode } from "./parser";
import { tokenizer } from "./tokenizer";

export function compiler(text:string):RootNode{
    return parser(tokenizer(text))
}