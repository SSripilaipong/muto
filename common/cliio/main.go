package cliio

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SSripilaipong/muto/common/rslt"
)

func ReadInputOneLine() rslt.Of[string] {
	return rslt.Fmap(func(s string) string {
		return strings.TrimRight(s, "\n")
	})(rslt.New(bufio.NewReader(os.Stdin).ReadString('\n')))
}

func PrintStringWithNewLine(s string) {
	fmt.Println(s)
}
