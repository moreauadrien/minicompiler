package ast

import "minicompiler/token"

type Attrib interface{}

type Program struct {
	Statements []Statement
}

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type ExpressionStatement struct {
	Token      *token.Token `json:"-"`
	Expression Expression   `json:"statement"`
}

type InitStatement struct {
	Token    *token.Token `json:"-"`
	Expr     Expression   `json:"expression"`
	Location string       `json:"location"`
}

type AssignStatement struct {
	Token *token.Token `json:"-"`
	Left  Identifier   `json:"left"`
	Right Expression   `json:"right"`
}

