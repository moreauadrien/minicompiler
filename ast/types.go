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

type BlockStatement struct {
	Token      *token.Token `json:"-"`
	Statements []Statement  `json:"statements"`
}

type IfStatement struct {
	Token       *token.Token    `json:"-"`
	Condition   Expression      `json:"condition"`
	Block       *BlockStatement `json:"block"`
	Alternative *BlockStatement `json:"alternative"`
}

type WhileStatement struct {
	Token       *token.Token    `json:"-"`
	Condition   Expression      `json:"condition"`
	Block       *BlockStatement `json:"block"`
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
