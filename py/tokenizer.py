# -*- coding: utf-8 -*-
# @Author: Jingyuexing
# @Date:   2024-05-05 09:23:29
# @Last Modified by:   Jingyuexing
# @Last Modified time: 2024-05-06 23:21:26
from enum import Enum

class TokenKind(Enum):
    GreatThan = "GreatThan"
    LessThan = "LessThan"
    LessThanOrEqual = "LessThanOrEqual"
    GreatThanOrEqual = "GreatThanOrEqual"
    COLON = "Colon"
    Number = "Number"
    Text = "Text"
    Identifier = "Identifier"
    NotEqual = "NotEqual"
    Comma = "Comma"
    Terminator = "Terminator"

class Token:
    def __init__(self, type:'TokenKind', value:str):
        self.type = type
        self.value = value

def create_token(type:"TokenKind", value:str):
    return Token(type, value)

def is_numeric(ch):
    return "0" <= ch <= "9"

def is_letter(ch):
    return ("a" <= ch <= "z") or ("A" <= ch <= "Z") or ch >= "\xff"

def is_whitespace(ch):
    return ch == "\n" or ch == " " or ch == "\r" or ch == "\t"

def tokenizer(text):
    current = 0
    tokens = []
    
    def advance():
        nonlocal current
        current += 1
    
    while current < len(text):
        ch = text[current]
        
        if is_whitespace(ch):
            while is_whitespace(ch) and current < len(text):
                advance()
                ch = text[current]
            tokens.append(create_token(TokenKind.Terminator, ";"))
            continue
        
        if is_numeric(ch):
            value = ""
            while (is_numeric(ch) or ch == ".") and current < len(text):
                value += ch
                advance()
                ch = text[current]
            tokens.append(create_token(TokenKind.Number, value))
        
        if is_letter(ch):
            value = ""
            while (is_letter(ch) or is_numeric(ch) or ch in ["-", "$", "_"]) and current < len(text):
                value += ch
                advance()
                ch = text[current]
            tokens.append(create_token(TokenKind.Identifier, value))
        
        if ch == ":":
            tokens.append(create_token(TokenKind.COLON, ch))
            advance()
            continue
        if ch == ",":
            tokens.append(create_token(TokenKind.Comma,ch))
            advance()
            continue
        if ch == "<":
            slice_ = text[current: current + 2]
            if slice_ == "<=":
                tokens.append(create_token(TokenKind.LessThanOrEqual, slice_))
                current += 2
            elif slice_ == "<>":
                tokens.append(create_token(TokenKind.NotEqual, slice_))
                current += 2
            else:
                tokens.append(create_token(TokenKind.LessThan, ch))
                advance()
        
        if ch == ">":
            slice_ = text[current: current + 2]
            if slice_ == ">=":
                tokens.append(create_token(TokenKind.GreatThanOrEqual, slice_))
                current += 2
            else:
                tokens.append(create_token(TokenKind.GreatThan, ch))
                advance()
    
    return tokens
