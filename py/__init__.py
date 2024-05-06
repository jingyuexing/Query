# -*- coding: utf-8 -*-
# @Author: Jingyuexing
# @Date:   2024-05-06 23:24:50
# @Last Modified by:   Jingyuexing
# @Last Modified time: 2024-05-06 23:39:39

from .tokenizer import tokenizer
from .parser import parse


def compiler(text:str):
    return parse(tokenizer(text))

if __name__ == '__main__':
    compiler("repo:name star:>3000 age:>400     version:>20")