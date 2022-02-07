package mif_parser

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var variable *regexp.Regexp = regexp.MustCompile("([a-z]{3,}_[a-z0-9]+)[ \t]+DCB[ \t]+0x([0-9A-F]+)")

var label *regexp.Regexp = regexp.MustCompile("^([a-z]{3,}[0-9]*)$")
var paramreg string = "#((?:0x[0-9A-F]+)|[a-z]{3,}_[a-z0-9]+)$"

var labelreg string = "([a-z]{3,}[0-9]*)$"
var opcode map[string]string = map[string]string{
	"MOV R1, #param":     "01",
	"MOV R0, #param":     "02",
	"MOV R3, #param":     "03",
	"LDRB R0, [R1]":      "04",
	"LDRB R3, [R1]":      "05",
	"STRB R0, [R1]":      "06",
	"CMP R0, R3":         "07",
	"BEQ label":          "08",
	"ADD R1, R1, R0":     "09",
	"BCC label":          "0A",
	"B label":            "0B",
	"LDRB R3, [R3]":      "0C",
	"ADD R0, R0, #param": "0D",
	"SUB R0, R0, R3":     "0E",
	"AND R0, R0, R3":     "0F",
	"LSR R0, R0, R3":     "10",
	"ADD R0, R0, R3":     "11",
	"ADD R0, R0, R0":     "12",
	"MOV R3, R0":         "13",
	"BNE label":          "14",
	"WAIT #param":        "15",
	"MUL R0, R0, R3":     "16",
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
	regmap := map[*regexp.Regexp]string{}
	for s, op := range opcode {
		s = strings.ReplaceAll(s, "[", `\[`)
		s = strings.ReplaceAll(s, "]", `\]`)
		s = strings.ReplaceAll(s, "#param", paramreg)
		s = strings.ReplaceAll(s, "label", labelreg)

		regmap[regexp.MustCompile(s)] = op
	}
	var b bytes.Buffer
	//entête fichier mif :
	entete := "DEPTH=256;\nWIDTH=24;\n\nADDRESS_RADIX=HEX;\nDATA_RADIX=HEX;\n\nCONTENT\nBEGIN\n"
	write(&b, "%v", entete)
	content := asmContent
	lines := strings.Split(content, "\n")
	lineMap := map[string]string{}
	i := 0
	for _, line := range lines {
		// mapping label et variable
		line = strings.Trim(line, " ")

		for reg, _ := range regmap {
			if reg.MatchString(line) {
				i++
			}
		}
		//remplissage de la map varline
		if g := variable.FindStringSubmatch(line); !isEmpty(line) && len(g) > 0 {
			lineMap[g[1]] = fmt.Sprintf("%4X", i)
			lineMap[g[1]] = strings.ReplaceAll(lineMap[g[1]], " ", "0")
			i++
			fmt.Println("key : ", g[1], "\t\tvalue : ", lineMap[g[1]])

		}
		//remplissage de la map labelLine
		if g := label.FindStringSubmatch(line); !isEmpty(line) && len(g) > 0 {

			lineMap[g[1]] = fmt.Sprintf("%4X", i)
			lineMap[g[1]] = strings.ReplaceAll(lineMap[g[1]], " ", "0")
			fmt.Println("key : ", g[1], "\t\tvalue : ", lineMap[g[1]])
		}

	}

	i = 0
	for _, line := range lines {

		line = strings.Trim(line, " ")

		//case of operation

		for reg, op := range regmap {

			if reg.MatchString(line) {
				g := reg.FindStringSubmatch(line)
				if len(g) > 1 {
					//numerotation ligne + opcode
					data := fmt.Sprintf("%2X", i)
					data = strings.ReplaceAll(data, " ", "0")
					data = fmt.Sprintf("%v : %v", data, op)
					//comentaire
					com := fmt.Sprintf("%v %v %v", "%", line, "%")
					//parmam

					var param string
					if strings.HasPrefix(g[1], "0x") {
						param = fmt.Sprintf("%4v", strings.TrimLeft(g[1], "0x"))
						param = strings.ReplaceAll(param, " ", "0")
					} else {
						param = fmt.Sprintf("%v", lineMap[g[1]])
						param = strings.ReplaceAll(param, " ", "0")
					}
					//ligne + opcode + param + commentaire
					data = fmt.Sprintf("%v%v;\t%v\n", data, param, com)
					write(&b, "%v", data)
					i++
				} else {
					// g := reg.FindStringSubmatch(line)
					//numerotation ligne + opcode
					data := fmt.Sprintf("%2X", i)
					data = strings.ReplaceAll(data, " ", "0")
					data = fmt.Sprintf("%v : %v", data, op)
					//comentaire
					com := fmt.Sprintf("%v %v %v", "%", line, "%")
					//parmam
					param := "0000"
					//ligne + opcode + param + commentaire
					data = fmt.Sprintf("%v%v;\t%v\n", data, param, com)
					write(&b, "%v", data)
					i++
				}
			}
		}

		//case of variable
		if g := variable.FindStringSubmatch(line); !isEmpty(line) && len(g) > 0 {
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
			data = fmt.Sprintf("%v%v;\t%v\n", data, value, com)

			write(&b, "%v", data)
			i++
		}

	}
	eof := "END"
	write(&b, "%v", eof)

	return b
}
