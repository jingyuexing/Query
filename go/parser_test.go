package query_test

import (
	"fmt"
	"testing"

	"github.com/jingyuexing/query"
)

func TestParser(t *testing.T) {
	tokens := query.Tokenizer("bigone age:>100 version:>1.3.0, tags:海,我,飞鸟")
	ast := query.Parse(tokens)
	if len(ast.Children) != 4 {
		t.Error("parser has error")
	}
}

func TestTransfomer(t *testing.T) {
	tokens := query.Tokenizer("bigone age:>100 version:>1.3.0, tags:海,我,飞鸟")
	ast := query.Parse(tokens)
	trans := query.Transfomer(ast)

	fmt.Printf("%s", trans)
}

func TestToken(t *testing.T) {
	tokens := query.Tokenizer("name age:>20")
	tokensLength := len(tokens)
	if tokensLength != 6 {
		t.Error("tokenizer has error")
	}
}
