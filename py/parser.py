# -*- coding: utf-8 -*-
# @Author: Jingyuexing
# @Date:   2024-05-05 09:23:16
# @Last Modified by:   Jingyuexing
# @Last Modified time: 2024-05-06 23:25:48
from enum import Enum
from typing import List
from tokenizer import Token, TokenKind

class ASTNodeKind(Enum):
    Root = "Query"
    Expression = "Expression"
    ConditionExpression = "ConditionExpression"
    Number = "Number"
    Text = "Text"


class Node:
    def __init__(self, type:'ASTNodeKind'):
        self._type = type

class ConditionNode(Node):
    _type:'ASTNodeKind'
    name:str
    condition:List['Literal']
    value:str
    def __init__(self, name:str, condition:List['Literal'], value:str):
        self._type = ASTNodeKind.ConditionExpression
        self.name = name
        self.condition = condition
        self.value = value

class RootNode(Node):
    def __init__(self):
        self._type:'ASTNodeKind' = ASTNodeKind.Root
        self.child:List['ConditionNode'] = []

def create_root_node():
    return RootNode()

class Literal(Node):
    def __init__(self, value:str, type:'ASTNodeKind'):
        self._type = type
        self.value = value

def create_condition_node(name, value):
    return ConditionNode(name, [], value)

def create_literal_node(type, value):
    return Literal(value, type)

def parse(tokens:List[Token]):
    current = 0
    token:Token = tokens[current]

    def eat():
        nonlocal current
        current += 1
        token = tokens[current]
        return token

    rootNode = create_root_node()

    while current < len(tokens):
        cache = []
        
        if token.type == TokenKind.Identifier:
            cache.append(token)
            token = eat()
        
        if token.type == TokenKind.COLON:
            token = eat()
            if token:
                condition = create_condition_node(cache[-1].value if cache else "", "=")
                while token and token.type != TokenKind.Terminator:
                    if token.type in [TokenKind.GreatThan, TokenKind.GreatThanOrEqual, TokenKind.LessThan, TokenKind.LessThanOrEqual, TokenKind.NotEqual]:
                        condition.value = token.value
                    elif token.type == TokenKind.Number:
                        condition.condition.append(create_literal_node(ASTNodeKind.Number, token.value))
                    elif token.type == TokenKind.Comma:
                        token = eat()
                        continue
                    elif token.type == TokenKind.Identifier:
                        condition.condition.append(create_literal_node(ASTNodeKind.Text, token.value))
                    token = eat()
                rootNode.child.append(condition)
        
        if token and token.type == TokenKind.Terminator:
            token = eat()
            continue
        current += 1

    return rootNode
