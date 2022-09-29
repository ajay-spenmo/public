package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const testFuncPattern = `func Test.*\(t \*testing\.T\) \{`

var r = regexp.MustCompile(testFuncPattern)

// RenameTestFunctions renames snake-cased test function names in source to pascal-cased names and returns
// the result.
func RenameTestFunctions(source string) string {
	var result string

	lines := strings.Split(source, "\n")
	for _, line := range lines {
		if !r.MatchString(line) {
			result += line + "\n"
			continue
		}

		funcName := strings.Split(strings.Split(line, "func Test")[1], "(")[0]
		var (
			modifiedFuncName  string
			underScorePrecede bool
		)
		for _, c := range funcName {
			if c != '_' {
				if underScorePrecede {
					modifiedFuncName += strings.ToUpper(string(c))
				} else {
					modifiedFuncName += string(c)
				}
			}
			underScorePrecede = c == '_'
		}
		result += fmt.Sprintf("func Test%s(t *testing.T) {\n", modifiedFuncName)
	}
	return result
}

func main() {
	source, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	result := RenameTestFunctions(string(source))
	fmt.Print(result)
}
