package ast

import "minicompiler/token"

type Attrib interface{}

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


type Program struct {
	Statements []Statement
}

func (p Program) TokenLiteral() string {
	return "Program"
}


type ExpressionStatement struct {
	Token      *token.Token `json:"-"`
	Expression Expression   `json:"statement"`
}

func (es ExpressionStatement) statementNode() {}
func (es ExpressionStatement) TokenLiteral() string {
	return "ExpressionStatement"
}


type InitStatement struct {
	Token    *token.Token `json:"-"`
	Expr     Expression   `json:"expression"`
	Location string       `json:"location"`
}

func (is InitStatement) statementNode() {}
func (is InitStatement) TokenLiteral() string {
	return "InitStatement"
}


type TabInitStatement struct {
	Token        *token.Token `json:"-"`
	Size         int          `json:"size"`
	DefaultValue int          `json:"defaultValue"`
	Location     string       `json:"location"`
}

func (oe TabInitStatement) statementNode() {}
func (oe TabInitStatement) TokenLiteral() string {
	return "TabInitStatement"
}


type AssignStatement struct {
	Token *token.Token `json:"-"`
	Left  Identifier   `json:"left"`
	Right Expression   `json:"right"`
}

func (ls AssignStatement) statementNode() {}
func (ls AssignStatement) TokenLiteral() string {
	return "AssignStatement"
}


type AssignTabStatement struct {
	Token *token.Token `json:"-"`
	Left  Identifier   `json:"left"`
	Index Expression `json:"index"`
	Right Expression   `json:"right"`
}

func (ls AssignTabStatement) statementNode() {}
func (ls AssignTabStatement) TokenLiteral() string {
	return "AssignTabStatement"
}


type BlockStatement struct {
	Token      *token.Token `json:"-"`
	Statements []Statement  `json:"statements"`
}

func (bs BlockStatement) statementNode() {}
func (bs BlockStatement) TokenLiteral() string {
	return "BlockStatement"
}


type IfStatement struct {
	Token       *token.Token    `json:"-"`
	Condition   Expression      `json:"condition"`
	Block       *BlockStatement `json:"block"`
	Alternative *BlockStatement `json:"alternative"`
}

func (is IfStatement) statementNode() {}
func (is IfStatement) TokenLiteral() string {
	return "IfStatement"
}


type WhileStatement struct {
	Token     *token.Token    `json:"-"`
	Condition Expression      `json:"condition"`
	Block     *BlockStatement `json:"block"`
}

func (is WhileStatement) statementNode() {}
func (is WhileStatement) TokenLiteral() string {
	return "WhileStatement"
}


type Identifier struct {
	Token *token.Token `json:"-"`
	Value string       `json:"value"`
}

func (i Identifier) expressionNode() {}
func (i Identifier) TokenLiteral() string {
	return string(i.Token.Lit)
}


type IntegerLiteral struct {
	Token *token.Token `json:"-"`
	Value string       `json:"value"`
}

func (il IntegerLiteral) expressionNode() {}
func (il IntegerLiteral) TokenLiteral() string {
	return string(il.Token.Lit)
}


type InfixExpression struct {
	Token    *token.Token `json:"-"`
	Type     string       `json:"-"`
	Left     Expression   `json:"left"`
	Right    Expression   `json:"right"`
	Operator string       `json:"operator"`
}

func (oe InfixExpression) expressionNode() {}
func (oe InfixExpression) TokenLiteral() string {
	return string(oe.Token.Lit)
}
