package evaluator

import (
	"github.com/ujuc/monkey/internal/ast"
	"github.com/ujuc/monkey/internal/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// 명령문
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// 표현식
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
