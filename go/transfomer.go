package query

import "strings"

func processExpression(ast ConditionNode) string {
	name := ast.Name
	finalString := make([]string, 0)
	if len(ast.Condition) > 0 {
		for _, exp := range ast.Condition {
			part := strings.Join([]string{name, exp.Operator, exp.Value.Value}, " ")
			finalString = append(finalString, part)
		}
	} else {
		finalString = append(finalString, name)
	}
	return "(" + strings.Join(finalString, " OR ") + ")"
}

func Transfomer(ast RootNode) string {
	expression := make([]string, 0)
	for _, condtion := range ast.Children {
		expression = append(expression, processExpression(condtion))
	}
	return strings.Join(expression, " AND ")
}
