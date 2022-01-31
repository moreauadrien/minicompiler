package mif_parser 

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var label *regexp.Regexp = regexp.MustCompile("([a-z]{3,}_[a-z0-9]+)[ \t]+DCB[ \t]+0x([0-9A-F]+)")

var opcode map[string]string = map[string]string{
	"MOV R1, ":       "01",
	"MOV R0, ":       "02",
	"MOV R3, ":       "03",
	"LDRB R0, [R1]":  "04",
	"LDRB R3, [R1]":  "05",
	"STRB  R0, [R1]": "06",
	"CMP  R0, R3":    "07",
	"BEQ ":           "08",
	"ADD R1, R1, R0": "09",
	"BCC ":           "0A",
	"B ":             "0B",
	"LDRB R3, [R3]":  "0C",
	"ADD R0, R0, ":   "0D",
	"SUB R0,R0, R3":  "0E",
	"AND R0,R0,R3":   "0F",
	"LSR R0, R0, R3": "10",
	"ADD R0, R0,R3":  "11",
	"ADD R0, R0,R0":  "12",
	"MOV R3, R0":     "13",
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isEmpty(line string) bool {
	line = strings.ReplaceAll(line, " ", "")
	if len(line) == 0 {
		return true
	} else {
		return false
	}
}

func write(b *bytes.Buffer, code string, args ...interface{}) {
	b.WriteString(fmt.Sprintf(code, args...))
}

func CompileToMif(asmContent string) bytes.Buffer {
	var b bytes.Buffer
	content := asmContent
	lines := strings.Split(content, "\n")
	i := 0
	for _, line := range lines {

		line = strings.Trim(line, " ")

		//case of variable
		if g := label.FindStringSubmatch(line); !isEmpty(line) && len(g) > 0 {
			//numerotation ligne
			data := fmt.Sprintf("%2X", i)
			data = strings.ReplaceAll(data, " ", "0")
			data = fmt.Sprintf("%v : ", data)
			//ecriture value
			value := fmt.Sprintf("%6v", g[2])
			value = strings.ReplaceAll(value, " ", "0")
			//comentaire
			com := fmt.Sprintf("%v %v %v", "%", line, "%")
			//ligne + value + commentaire
			data = fmt.Sprintf("%v%v\t%v\n", data, value, com)

			write(&b, "%v", data)
			i++
		}

	}

	return b 
}
