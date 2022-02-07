package gen

import (
	"bytes"
	"fmt"
	"minicompiler/ast"
	"strconv"
	"strings"
)

var tmpCount int
var labelCount int

var operatorToInstru = map[string]string{
	"+":  "ADD",
	"-":  "SUB",
	"*":  "MUL",
	"==": "BEQ",
	"<":  "BCC",
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func write(b *bytes.Buffer, code string, args ...interface{}) {
	b.WriteString(fmt.Sprintf(code, args...))
}

func GenWrapper(p *ast.Program) bytes.Buffer {
	tmpCount = 0
	labelCount = 0
	var b, bVar, bTempVar, bTabs bytes.Buffer
	gen(p, &b, &bVar, &bTempVar, &bTabs)

	b.WriteString("endprog\nB endprog\n\n")

	b.WriteString("const_1 DCB 0x1\n")
	b.WriteString("const_0 DCB 0x0\n")
	b.WriteString(bTempVar.String())
	b.WriteString(bTabs.String())
	b.WriteString(bVar.String())

	return b
}

func newTempVariable(bTempVar *bytes.Buffer, value string) string {
	tmpCount++
	varName := fmt.Sprintf("temp_%v", tmpCount)
	write(bTempVar, "%v DCB %v\n", varName, value)

	return varName
}

func newLabelNumber() int {
	labelCount++
	return labelCount
}

func gen(node ast.Node, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	switch node := node.(type) {

	case *ast.Program:
		return genProgram(node, b, bVar, bTempVar, bTabs)
	case *ast.ExpressionStatement:
		return genExpressionStatement(node, b, bVar, bTempVar, bTabs)
	case *ast.AssignStatement:
		return genAssignStatement(node, b, bVar, bTempVar, bTabs)
	case *ast.AssignTabStatement:
		return genAssignTabStatement(node, b, bVar, bTempVar, bTabs)
	case *ast.InitStatement:
		return genInitStatement(node, b, bVar, bTempVar, bTabs)
	case *ast.BlockStatement:
		return genBlockStatement(node, b, bVar, bTempVar, bTabs)
	case *ast.InfixExpression:
		return genInfixExpression(node, b, bVar, bTempVar, bTabs)
	case *ast.IntegerLiteral:
		return genInteger(node, b, bVar, bTempVar, bTabs)
	case *ast.Identifier:
		return genIdentifier(node, b, bVar, bTempVar, bTabs)
	case *ast.IfStatement:
		return genIfStatement(node, b, bVar, bTempVar, bTabs)
	case *ast.WhileStatement:
		return genWhileStatement(node, b, bVar, bTempVar, bTabs)
	case *ast.TabInitStatement:
		return genTabInitStatement(node, b, bVar, bTempVar, bTabs)
	}
	return ""
}

func genProgram(node *ast.Program, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	for _, stmt := range node.Statements {
		gen(stmt, b, bVar, bTempVar, bTabs)
	}
	return ""
}

func genExpressionStatement(node *ast.ExpressionStatement, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	value := gen(node.Expression, b, bVar, bTempVar, bTabs)
	write(b, "%v\n", value)
	return ""
}

func genAssignStatement(node *ast.AssignStatement, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	right := gen(node.Right, b, bVar, bTempVar, bTabs)
	write(b, "MOV R1, #%v\n", right)
	write(b, "LDRB R0, [R1]\n")
	write(b, "MOV R1, #var_%v\n", node.Left.Value)
	write(b, "STRB R0, [R1]\n\n")
	return ""
}

func genAssignTabStatement(node *ast.AssignTabStatement, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	index := gen(node.Index, b, bVar, bTempVar, bTabs)
	right := gen(node.Right, b, bVar, bTempVar, bTabs)
	ident := node.Left.Value

	write(b, "MOV R1, #%v\n", right)
	write(b, "LDRB R3, [R1]\n")
	write(b, "MOV R1, #%v\n", index)
	write(b, "LDRB R0, [R1]\n")
	if ident == "screen" {
		write(b, "MOV R1, #0x200\n")
	} else {
		write(b, "MOV R1, #tab_%v\n", node.Left.Value)
	}
	write(b, "ADD R1, R1, R0\n")
	write(b, "STRB R3, [R1]\n")
	return ""
}

func genInitStatement(node *ast.InitStatement, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	right := gen(node.Expr, b, bVar, bTempVar, bTabs)
	write(b, "MOV R1, #%v\n", right)
	write(b, "LDRB R0, [R1]\n")
	write(b, "MOV R1, #var_%v\n", node.Location)
	write(b, "STRB R0, [R1]\n\n")

	write(bVar, "var_%v DCB 0x0\n", node.Location)
	return ""
}

func genTabInitStatement(node *ast.TabInitStatement, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	size := node.Size
	defaultValue := fmt.Sprintf("0x%X", node.DefaultValue)

	tabInit := strings.Repeat(defaultValue+",", size)
	tabInit = strings.TrimRight(tabInit, ", ")

	write(bTabs, "tab_%v DCB %v\n", node.Location, tabInit)

	return ""
}

func genInteger(node *ast.IntegerLiteral, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	i, err := strconv.ParseUint(node.Value, 10, 32)
	check(err)
	hex := fmt.Sprintf("0x%X", i)
	tmp := newTempVariable(bTempVar, hex)

	return tmp
}

func genIdentifier(node *ast.Identifier, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	return "var_" + node.Value
}

func genInfixExpression(node *ast.InfixExpression, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	tempLabel := newLabelNumber()
	left := gen(node.Left, b, bVar, bTempVar, bTabs)
	right := gen(node.Right, b, bVar, bTempVar, bTabs)

	write(b, "MOV R1, #%v\n", left)
	write(b, "LDRB R0, [R1]\n")
	write(b, "MOV R1, #%v\n", right)
	write(b, "LDRB R3, [R1]\n")

	switch node.Operator {
	case "+", "-", "*":
		write(b, "%v R0, R0, R3\n", operatorToInstru[node.Operator])

	case "==", "<":
		write(b, "CMP R0, R3\n")
		write(b, "MOV R1, #const_1\n")
		write(b, "LDRB R0, [R1]\n")
		write(b, "%v condtrue%v\n", operatorToInstru[node.Operator], tempLabel)
		write(b, "MOV R1, #const_0\n")
		write(b, "LDRB R0, [R1]\n")
		write(b, "condtrue%v\n", tempLabel)
	}

	tmp := newTempVariable(bTempVar, "0x0")

	write(b, "MOV R1, #%v\n", tmp)
	write(b, "STRB R0, [R1]\n\n")

	return tmp
}

func genIfStatement(node *ast.IfStatement, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	cond := gen(node.Condition, b, bVar, bTempVar, bTabs)

	labelId := newLabelNumber()

	elseCode := gen(node.Alternative, b, bVar, bTempVar, bTabs)

	write(b, "MOV R1, #%v\n", cond)
	write(b, "LDRB R0, [R1]\n")
	write(b, "MOV R3, #0x1\n")
	write(b, "CMP R0, R3\n")
	if len(elseCode) > 0 {
		write(b, "BNE else%v\n", labelId)
	}
	gen(node.Block, b, bVar, bTempVar, bTabs)
	write(b, "B ifend%v\n", labelId)

	if len(elseCode) > 0 {
		write(b, "else%v\n", labelId)
		gen(node.Alternative, b, bVar, bTempVar, bTabs)
	}
	write(b, "ifend%v\n", labelId)
	return ""
}

func genWhileStatement(node *ast.WhileStatement, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	labelId := newLabelNumber()
	write(b, "startwhile%v\n", labelId)
	cond := gen(node.Condition, b, bVar, bTempVar, bTabs)
	write(b, "MOV R1, #%v\n", cond)
	write(b, "LDRB R0, [R1]\n")
	write(b, "MOV R3, #0x1\n")
	write(b, "CMP R0, R3\n")
	write(b, "BNE endwhile%v\n", labelId)
	gen(node.Block, b, bVar, bTempVar, bTabs)
	write(b, "B startwhile%v\n", labelId)
	write(b, "endwhile%v\n", labelId)

	return ""
}

func genBlockStatement(node *ast.BlockStatement, b, bVar, bTempVar, bTabs *bytes.Buffer) string {
	for _, stmt := range node.Statements {
		gen(stmt, b, bVar, bTempVar, bTabs)
	}
	return ""
}
