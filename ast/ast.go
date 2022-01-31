package ast

import (
	"errors"
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


func NewProgram(stmts Attrib) (*Program, error) {
	s, ok := stmts.([]Statement)
	if !ok {
		return nil, errors.New(fmt.Sprintf("NewProgram []Statement stmts %v", stmts))
	}

	return &Program{Statements: s}, nil
}

func NewStatementList() ([]Statement, error) {
	return []Statement{}, nil
}

func NewExpressionStatement(expr Attrib) (Statement, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, errors.New(fmt.Sprintf("NewExpressionStatement Expression expr %s", expr))
	}
	return &ExpressionStatement{Expression: e}, nil
}

func AppendStatement(stmtList, stmt Attrib) ([]Statement, error) {
	s, ok := stmt.(Statement)
	if !ok {
		return nil, errors.New(fmt.Sprintf("AppendStatement Statement stmt %s", stmt))
	}
	return append(stmtList.([]Statement), s), nil
}

func NewAssignStatement(left, right Attrib) (Statement, error) {
	l, ok := left.(*token.Token)
	if !ok {
		return nil, errors.New(fmt.Sprintf("NewAssignStatement Identifier left %v", left))
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, errors.New(fmt.Sprintf("NewAssignStatement Identifier right %v", right))
	}

	return &AssignStatement{Left: Identifier{Value: string(l.Lit)}, Right: r}, nil
}


type Identifier struct {
	Token *token.Token `json:"-"`
	Value string       `json:"value"`
}

type IntegerLiteral struct {
	Token *token.Token `json:"-"`
	Value string       `json:"value"`
}

type InfixExpression struct {
	Token    *token.Token `json:"-"`
	Type     string       `json:"-"`
	Left     Expression   `json:"left"`
	Right    Expression   `json:"right"`
	Operator string       `json:"operator"`
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


func NewInfixExpression(left, right, oper Attrib) (Expression, error) {
	l, ok := left.(Expression)
	if !ok {
		fmt.Println(left)
		return nil, errors.New(fmt.Sprintf("NewInfixExpression Expression left %v", left))
	}

	o, ok := oper.(*token.Token)
	if !ok {
		return nil, errors.New(fmt.Sprintf("NewInfixExpression *token.Token oper %v", oper))
	}

	r, ok := right.(Expression)
	if !ok {
		return nil, errors.New(fmt.Sprintf("NewInfixExpression Expression right %v", right))
	}

	return &InfixExpression{Left: l, Operator: string(o.Lit), Right: r, Token: o}, nil
}

func NewIntegerLiteral(integer Attrib) (Expression, error) {
	intLit, ok := integer.(*token.Token)
	if !ok {
		return nil, errors.New(fmt.Sprintf("NewIntegerLiteral *token.Token integer %v", integer))
	}

	return &IntegerLiteral{Token: intLit, Value: string(intLit.Lit)}, nil
}

func NewIdentInit(ident, expr Attrib) (Statement, error) {
	e, ok := expr.(Expression)
	if !ok {
		return nil, errors.New(fmt.Sprintf("NewIdentInit Expression expr %v", expr))
	}

	return &InitStatement{Location: string(ident.(*token.Token).Lit), Token: ident.(*token.Token), Expr: e}, nil
}

func NewIdentExpression(ident Attrib) (*Identifier, error) {
	return &Identifier{Value: string(ident.(*token.Token).Lit), Token: ident.(*token.Token)}, nil
}
