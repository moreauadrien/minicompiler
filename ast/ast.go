package ast

import (
	"fmt"
	"minicompiler/token"
	"strconv"
)

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

func NewAssignTabStatement(left, index, right Attrib) (Statement, error) {
	l, ok := left.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("NewAssignTabStatement Identifier left %v", left)
	}

	i, ok := index.(Expression)
	if !ok {
		return nil, fmt.Errorf("NewAssignTabStatement Identifier index %v", index)
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, fmt.Errorf("NewAssignTabStatement Expression right %v", right)
	}

	return &AssignTabStatement{Left: Identifier{Value: string(l.Lit)}, Right: r, Index: i}, nil
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

func NewTabInit(ident, size, defaultValue Attrib) (Statement, error) {
	tabSize, err := strconv.Atoi(string(size.(*token.Token).Lit))

	if err != nil {
		return nil, fmt.Errorf("NewTabInit size %v", size)
	}

	defaultVal, err := strconv.Atoi(string(defaultValue.(*token.Token).Lit))

	if err != nil {
		return nil, fmt.Errorf("NewTabInit defaultValue %v", defaultValue)
	}

	return &TabInitStatement{Token: ident.(*token.Token), Location: string(ident.(*token.Token).Lit), Size: tabSize, DefaultValue: defaultVal}, nil
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

func NewWaitStatement(time Attrib) (Statement, error) {
	timeInt, err := strconv.Atoi(string(time.(*token.Token).Lit))

	if err != nil {
		return nil, fmt.Errorf("NewWaitStatement timeInt %v", timeInt)
	}

	return &WaitStatement{Token: time.(*token.Token), Time: timeInt}, nil
}
