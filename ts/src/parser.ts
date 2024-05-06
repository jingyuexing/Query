import { type Token, TokenKind, tokenizer } from "./tokenizer"

export enum ASTNodeKind {
    Root = "Query",
    Expression = "Expression",
    ConditionExpression = "ConditionExpression",
    Number = "Number",
    Text = "Text"
}

interface Node {
    type: ASTNodeKind
}

interface ConditionNode extends Node {
    // name 表达式类型
    name: string
    // 表达式条件
    condition: Literal[]
    // 表达式原始符号
    value: string
}

export interface RootNode extends Node {
    child:ConditionNode[]
}

function createRootNode():RootNode{
    return {
        type:ASTNodeKind.Root,
        child:[]
    }
}

interface Literal extends Node {
    value:string
}

function createConditionNode(name: string, value: string): ConditionNode {
    return {
        type: ASTNodeKind.ConditionExpression,
        value,
        condition: [],
        name
    }
}
function createLiteralNode(type: ASTNodeKind, value: string):Literal{
    return {
        type,
        value
    }
}
export function parser(tokens: Token[]) {
    let current = 0
    let token = tokens[current]
    function eat() {
        current++
        token = tokens[current]
    }
    const rootNode = createRootNode()
    while (current < tokens.length) {
        let cache:Token[] = []
        if(token.type === TokenKind.Identifier){
            cache.push(token)
            eat()
        }
        if(token.type === TokenKind.COLON){
            current++
            token = tokens[current]
            const condition = createConditionNode(cache.pop()?.value || "","=")
            while(current < tokens.length && token.type !== TokenKind.Terminator){
                switch(token.type){
                    case TokenKind.GreatThan:
                    case TokenKind.GreatThanOrEqual:
                    case TokenKind.LessThan:
                    case TokenKind.LessThanOrEqual:
                    case TokenKind.NotEqual:
                        condition.value = token.value
                        break;
                    case TokenKind.Comma:
                        eat()
                        continue;
                    case TokenKind.Number:
                        condition.condition.push(createLiteralNode(ASTNodeKind.Number,token.value))
                        break;
                    case TokenKind.Identifier:
                        condition.condition.push(createLiteralNode(ASTNodeKind.Text,token.value))
                        break;
                }
                eat()
            }
            rootNode.child.push(condition)
        }
        if(token && token.type === TokenKind.Terminator){
            eat()
            continue
        }
        current++
    }
    return rootNode
}