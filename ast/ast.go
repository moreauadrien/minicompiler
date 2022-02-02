package ast

import (
	"fmt"
	"minicompiler/token"
)

// interface methods
func (p Program) TokenLiteral() string {
	return "Program"
}

// Statements
func (is InitStatement) statementNode() {}
func (is InitStatement) TokenLiteral() string {
	return "InitStatement"
}

func (es ExpressionStatement) statementNode() {}
func (es ExpressionStatement) TokenLiteral() string {
	return "ExpressionStatement"
}

func (ls AssignStatement) statementNode() {}
func (ls AssignStatement) TokenLiteral() string {
	return "AssignStatement"
}

func (is IfStatement) statementNode() {}
func (is IfStatement) TokenLiteral() string {
	return "IfStatement"
}

func (is WhileStatement) statementNode() {}
func (is WhileStatement) TokenLiteral() string {
	return "WhileStatement"
}

func (bs BlockStatement) statementNode() {}
func (bs BlockStatement) TokenLiteral() string {
	return "BlockStatement"
}

func (i Identifier) expressionNode() {}
func (i Identifier) TokenLiteral() string {
	return string(i.Token.Lit)
}

func (il IntegerLiteral) expressionNode() {}
func (il IntegerLiteral) TokenLiteral() string {
	return string(il.Token.Lit)
}

func (oe InfixExpression) expressionNode() {}
func (oe InfixExpression) TokenLiteral() string {
	return string(oe.Token.Lit)
}

func NewProgram(stmts Attrib) (*Program, error) {
	s, ok := stmts.([]Statement)
	if !ok {
		return nil, fmt.Errorf("NewProgram []Statement stmts %v", stmts)
	}

	return &Program{Statements: s}, nil
}

func NewStatementList() ([]Statement, error) {
	return []Statement{}, nil
}

func NewExpressionStatement(expr Attrib) (Statement, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, fmt.Errorf("NewExpressionStatement Expression expr %s", expr)
	}
	return &ExpressionStatement{Expression: e}, nil
}

func AppendStatement(stmtList, stmt Attrib) ([]Statement, error) {
	s, ok := stmt.(Statement)
	if !ok {
		return nil, fmt.Errorf("AppendStatement Statement stmt %s", stmt)
	}
	return append(stmtList.([]Statement), s), nil
}

func NewAssignStatement(left, right Attrib) (Statement, error) {
	l, ok := left.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("NewAssignStatement Identifier left %v", left)
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, fmt.Errorf("NewAssignStatement Identifier right %v", right)
	}

	return &AssignStatement{Left: Identifier{Value: string(l.Lit)}, Right: r}, nil
}

func NewBlockStatement(stmts Attrib) (*BlockStatement, error) {
	s, ok := stmts.([]Statement)
	if !ok {
		return nil, fmt.Errorf("NewBlockStatement []Statement stmts %v", stmts)
	}

	return &BlockStatement{Statements: s}, nil
}

func NewInfixExpression(left, right, oper Attrib) (Expression, error) {
	l, ok := left.(Expression)
	if !ok {
		fmt.Println(left)
		return nil, fmt.Errorf("NewInfixExpression Expression left %v", left)
	}

	o, ok := oper.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("NewInfixExpression *token.Token oper %v", oper)
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, fmt.Errorf("NewInfixExpression Expression right %v", right)
	}

	return &InfixExpression{Left: l, Operator: string(o.Lit), Right: r, Token: o}, nil
}

func NewIntegerLiteral(integer Attrib) (Expression, error) {
	intLit, ok := integer.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("NewIntegerLiteral *token.Token integer %v", integer)
	}

	return &IntegerLiteral{Token: intLit, Value: string(intLit.Lit)}, nil
}

func NewIdentInit(ident, expr Attrib) (Statement, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, fmt.Errorf("NewIdentInit Expression expr %v", expr)
	}

	return &InitStatement{Location: string(ident.(*token.Token).Lit), Token: ident.(*token.Token), Expr: e}, nil
}

func NewIdentExpression(ident Attrib) (*Identifier, error) {
	return &Identifier{Value: string(ident.(*token.Token).Lit), Token: ident.(*token.Token)}, nil
}


func NewIfStatement(cond, cons, alt Attrib) (Statement, error) {
	c, ok := cond.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type of cond. got=%T", cond)
	}

	cs, ok := cons.(*BlockStatement)
	if !ok {
		return nil, fmt.Errorf("invalid type of cons. got=%T", cons)
	}

	a, ok := alt.(*BlockStatement)
	if !ok {
		return nil, fmt.Errorf("invalid type of alt. got=%T", alt)
	}

	return &IfStatement{Condition: c, Block: cs, Alternative: a}, nil
}

func NewWhileStatement(cond, cons Attrib) (Statement, error) {
	c, ok := cond.(Expression)
	if !ok {
		return nil, fmt.Errorf("invalid type of cond. got=%T", cond)
	}

	cs, ok := cons.(*BlockStatement)
	if !ok {
		return nil, fmt.Errorf("invalid type of cons. got=%T", cons)
	}

	return &WhileStatement{Condition: c, Block: cs}, nil
}
