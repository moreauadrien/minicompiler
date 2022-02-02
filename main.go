package main

import (
	"minicompiler/ast"
	"minicompiler/cmd"
	"minicompiler/lexer"
	"minicompiler/mif_parser"
	"minicompiler/parser"
	"os"
	"regexp"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filepath string) (string, error) {
	buffer, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func writeFile(filepath, content string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

func Parse(input string) *ast.Program {
	l := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	node, err := p.Parse(l)
	checkError(err)
	program, _ := node.(*ast.Program)

	return program
}

func main() {
	cmpOptions, err := cmd.GetCompileOptions()
	checkError(err)

	input, err := readFile(cmpOptions.Inputpath)
	checkError(err)
	//program := Parse(input)
	reg := regexp.MustCompile(`\..*?$`)
	mifFileName := reg.ReplaceAllString(cmpOptions.Inputpath, ".mif")

	//asmCode := gen.GenWrapper(program)

	//if cmpOptions.AssemblyOutput {
	//writeFile(cmpOptions.Outputpath, asmCode.String())

	mifCode := mif_parser.CompileToMif(input)
	writeFile(mifFileName, mifCode.String())
	//}

	/*	js, err := json.MarshalIndent(program, "", "    ")
		checkError(err)
		fmt.Printf("\n%s\n", js)*/
}
