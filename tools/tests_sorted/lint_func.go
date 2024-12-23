package main

import (
	"go/ast"
	"go/token"
	"slices"
)

func lintFuncDecl(funcSpec *ast.FuncDecl, fileSet *token.FileSet, issues *Issues) {
	lintFunc(funcSpec.Pos(), funcSpec.Type, funcSpec.Body, fileSet, issues)
}

func lintFuncLit(funcSpec *ast.FuncLit, fileSet *token.FileSet, issues *Issues) {
	lintFunc(funcSpec.Pos(), funcSpec.Type, funcSpec.Body, fileSet, issues)
}

func lintFunc(pos token.Pos, funcType *ast.FuncType, funcBody *ast.BlockStmt, fileSet *token.FileSet, issues *Issues) {
	if !isTestFunction(funcType) {
		return
	}

	subtests := topLevelRunNames(funcBody.List)
	sortedSubtests := make([]string, len(subtests))
	copy(sortedSubtests, subtests)
	slices.Sort(sortedSubtests)

	if slices.Equal(subtests, sortedSubtests) {
		return
	}
	*issues = append(*issues, issue{
		expected: sortedSubtests,
		position: fileSet.Position(pos),
	})
}

// isTestFunction returns true if a given funcSpec describes a test function/helper.
func isTestFunction(funcType *ast.FuncType) bool {
	params := funcType.Params.List
	// Find the first param.
	if len(params) < 1 {
		return false
	}
	first := params[0]
	// Check if it's `t *testing.T`.
	if len(first.Names) != 1 {
		// This is a weird func(a, b *testing.T) function.
		// Skip it.
		return false
	}
	firstName := first.Names[0].Name
	// Check first parameter name.
	if firstName != "t" {
		// Not a typical test function/helper.
		return false
	}
	if !isTestingT(first.Type) {
		return false
	}
	return true
}

func isTestingT(typeSpec ast.Expr) bool {
	ptr, ok := typeSpec.(*ast.StarExpr)
	if !ok {
		// Not a pointer.
		return false
	}
	selector, ok := ptr.X.(*ast.SelectorExpr)
	if !ok {
		// Not a `x.y` type.
		return false
	}
	selectorName := selector.Sel.Name
	selectorNamespace, ok := selector.X.(*ast.Ident)
	if !ok {
		// The left part of the selector isn't a simple identifier.
		return false
	}
	return selectorNamespace.Name == "testing" && selectorName == "T"
}

func topLevelRunNames(statements []ast.Stmt) []string {
	var subtests []string
	for _, stmt := range statements {
		expr, ok := stmt.(*ast.ExprStmt)
		if !ok {
			continue
		}
		call, ok := expr.X.(*ast.CallExpr)
		if !ok {
			continue
		}
		if !isTRun(call.Fun) {
			continue
		}
		// Get the first arg.
		if len(call.Args) < 1 {
			continue
		}
		firstArg := call.Args[0]
		lit, ok := firstArg.(*ast.BasicLit)
		if !ok {
			// It is meaningless to sort non-literals.
			continue
		}
		if lit.Kind != token.STRING {
			// Only strings can be sorted.
			continue
		}
		// FIXME: unescape
		subtests = append(subtests, lit.Value)
	}
	return subtests
}

// isTRun returns true if a given expression is `t.Run`.
func isTRun(expr ast.Expr) bool {
	selector, ok := expr.(*ast.SelectorExpr)
	if !ok {
		// Not a `a.b` expression.
		return false
	}
	selectorName := selector.Sel.Name
	selectorNamespace, ok := selector.X.(*ast.Ident)
	if !ok {
		// The left part of the selector isn't a simple identifier.
		return false
	}
	return selectorNamespace.Name == "t" && selectorName == "Run"
}
