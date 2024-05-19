package query

// ASTNodeKind 定义不同类型的 AST 节点
type ASTNodeKind string

const (
	Root                ASTNodeKind = "Query"
	Expression          ASTNodeKind = "Expression"
	ConditionExpression ASTNodeKind = "ConditionExpression"
	Number              ASTNodeKind = "Number"
	Text                ASTNodeKind = "Text"
)

// Node 表示一个 AST 节点
type Node struct {
	Type ASTNodeKind
}

// ConditionNode 表示一个条件表达式节点
type ConditionNode struct {
	Node
	Name      string
	Condition []ExpressionNode
	Value     string
}

type ExpressionNode struct {
	Node
	Value    Literal
	Operator string
}

// RootNode 表示根节点
type RootNode struct {
	Node
	Children []ConditionNode
}

// Literal 表示一个文字节点
type Literal struct {
	Node
	Value string
}

// createRootNode 创建根节点
func createRootNode() RootNode {
	return RootNode{Node: Node{Type: Root}, Children: []ConditionNode{}}
}

func createExpressionNode(operator string) ExpressionNode {
	return ExpressionNode{Node: Node{Type: Expression}, Operator: operator}
}

// createConditionNode 创建条件节点
func createConditionNode(name, value string) ConditionNode {
	return ConditionNode{
		Node:      Node{Type: ConditionExpression},
		Name:      name,
		Condition: []ExpressionNode{},
		Value:     value,
	}
}

// createLiteralNode 创建文字节点
func createLiteralNode(nodeType ASTNodeKind, value string) Literal {
	return Literal{
		Node:  Node{Type: nodeType},
		Value: value,
	}
}

// parse 解析 Token 列表并生成 AST
func Parse(tokens []Token) RootNode {
	current := 0
	token := tokens[current]

	eat := func() Token {
		if current < len(tokens) {
			current++
		}
		if current < len(tokens) {
			return tokens[current]
		}
		return Token{Type: ""}
	}

	rootNode := createRootNode()

	for current < len(tokens) {
		var cache []Token

		if token.Type == Identifier {
			cache = append(cache, token)
			token = eat()
		}

		if token.Type == COLON {
			token = eat()
			if token.Type != "" {
				var condition ConditionNode
				var cahche_exp = make([]ExpressionNode, 0)
				if len(cache) > 0 {
					condition = createConditionNode(cache[len(cache)-1].Value, "")
				} else {
					condition = createConditionNode("", "")
				}
				for token.Type != Terminator && current < len(tokens) {
					switch token.Type {
					case GreatThan, GreatThanOrEqual, LessThan, LessThanOrEqual, NotEqual:
						cahche_exp = append(cahche_exp, createExpressionNode(token.Value))
						token = eat()
						continue
					case TNumber, TText, Identifier:
						var expression ExpressionNode
						if len(cahche_exp) == 0 {
							expression = createExpressionNode("=")
						} else {
							expression = cahche_exp[len(cahche_exp)-1]
						}
						if token.Type == Identifier {
							token.Type = TText
						}
						var literal Literal
						if token.Type == TNumber {
							literal = createLiteralNode(Number, token.Value)
						} else {
							literal = createLiteralNode(Text, token.Value)
						}
						expression.Value = literal
						condition.Condition = append(condition.Condition, expression)
					case Comma:
						token = eat()
						continue
					case Quote:
						var expression ExpressionNode
						if len(cahche_exp) == 0 {
							expression = createExpressionNode("=")
						} else {
							expression = cahche_exp[len(cahche_exp)-1]
						}
						expression.Value = createLiteralNode(Text, token.Value)
						condition.Condition = append(condition.Condition, expression)
					}

					token = eat()
				}
				cache = make([]Token, 0)
				rootNode.Children = append(rootNode.Children, condition)
			}
		}

		if token.Type == Terminator {
			if len(cache) > 0 {
				condition := createConditionNode(cache[len(cache)-1].Value, "=")
				rootNode.Children = append(rootNode.Children, condition)
			}
			token = eat()
			continue
		}
		current++
	}

	return rootNode
}
