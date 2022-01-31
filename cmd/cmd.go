package cmd

import (
	"errors"
	"flag"
	"os"
	"regexp"
)

type CompileOptions struct {
	Inputpath      string
	Outputpath     string
	AssemblyOutput bool
}

func GetCompileOptions() (CompileOptions, error) {
	if len(os.Args) < 2 {
		return CompileOptions{}, errors.New("you must specify the path of the input file")
	}
	flag.Parse()

	inputPath := flag.Arg(0)
	outputPath := "" 

	if len(outputPath) == 0 {
		reg := regexp.MustCompile(`\..*?$`)
		outputPath = reg.ReplaceAllString(inputPath, ".asm")
	}

	return CompileOptions{inputPath, outputPath, true}, nil

}
