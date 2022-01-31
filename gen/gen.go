package gen

import (
	"bytes"
	"fmt"
	"minicompiler/ast"
	"strconv"
)

var tmpCount int

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
	var b, bVar, bTempVar bytes.Buffer
	gen(p, &b, &bVar, &bTempVar)

	b.WriteString("endprog END\n\n")

	b.WriteString(bTempVar.String())
	b.WriteString(bVar.String())

	return b
}

func newTempVariable(bTempVar *bytes.Buffer, value string) string {
	tmpCount++
	varName := fmt.Sprintf("temp_%v", tmpCount)
	write(bTempVar, "%v DCB %v\n", varName, value)

	return varName
}

func gen(node ast.Node, b, bVar, bTempVar *bytes.Buffer) string {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return genProgram(node, b, bVar, bTempVar)
	case *ast.ExpressionStatement:
		return genExpressionStatement(node, b, bVar, bTempVar)
	case *ast.AssignStatement:
		return genAssignStatement(node, b, bVar, bTempVar)
	case *ast.InitStatement:
		return genInitStatement(node, b, bVar, bTempVar)

	// // Expressions
	case *ast.InfixExpression:
		return genInfixExpression(node, b, bVar, bTempVar)
	case *ast.IntegerLiteral:
		return genInteger(node, b, bVar, bTempVar)
	case *ast.Identifier:
		return genIdentifier(node, b, bVar, bTempVar)
	}
	return ""
}

func genProgram(node *ast.Program, b, bVar, bTempVar *bytes.Buffer) string {
	for _, stmt := range node.Statements {
		gen(stmt, b, bVar, bTempVar)
	}
	return ""
}

func genExpressionStatement(node *ast.ExpressionStatement, b, bVar, bTempVar *bytes.Buffer) string {
	value := gen(node.Expression, b, bVar, bTempVar)
	write(b, "%v\n", value)
	return ""
}

func genAssignStatement(node *ast.AssignStatement, b, bVar, bTempVar *bytes.Buffer) string {
	right := gen(node.Right, b, bVar, bTempVar)
	write(b, "MOV R1, #%v\n", right)
	write(b, "LDRB R0, [R1]\n")
	write(b, "MOV R1, #var_%v\n", node.Left.Value)
	write(b, "STRB R0, [R1]\n\n")
	return ""
}

func genInitStatement(node *ast.InitStatement, b, bVar, bTempVar *bytes.Buffer) string {
	right := gen(node.Expr, b, bVar, bTempVar)
	write(b, "MOV R1, #%v\n", right)
	write(b, "LDRB R0, [R1]\n")
	write(b, "MOV R1, #var_%v\n", node.Location)
	write(b, "STRB R0, [R1]\n\n")

	write(bVar, "var_%v DCB 0x0\n", node.Location)
	return ""
}

func genInteger(node *ast.IntegerLiteral, b, bVar, bTempVar *bytes.Buffer) string {
	i, err := strconv.ParseUint(node.Value, 10, 32)
	check(err)
	hex := fmt.Sprintf("0x%X", i) 
	tmp := newTempVariable(bTempVar, hex)

	return tmp 
}

func genIdentifier(node *ast.Identifier, b, bVar, bTempVar *bytes.Buffer) string {
	return "var_" + node.Value
}

func genInfixExpression(node *ast.InfixExpression, b, bVar, bTempVar *bytes.Buffer) string {
	left := gen(node.Left, b, bVar, bTempVar)
	right := gen(node.Right, b, bVar, bTempVar)

	write(b, "MOV R1, #%v\n", left)
	write(b, "LDRB R0, [R1]\n")
	write(b, "MOV R1, #%v\n", right)
	write(b, "LDRB R3, [R1]\n")

	switch node.Operator {
	case "+":
		write(b, "ADD R0, R0, R3\n")

	case "-":
		write(b, "SUB R0, R0, R3\n")
	}

	tmp := newTempVariable(bTempVar, "0x0")

	write(b, "MOV R1, #%v\n", tmp)
	write(b, "STRB R0, [R1]\n\n")

	return tmp
}
