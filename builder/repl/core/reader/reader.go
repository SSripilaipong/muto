package reader

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/muto/builder/repl/core/command"
	"github.com/SSripilaipong/muto/common/optional"
)

type Reader struct {
	lineReader lineReader
}

func New() Reader {
	return Reader{lineReader: newLineReader()}
}

func (r Reader) Read() optional.Of[command.Command] {
	text, err := r.lineReader.ReadLine()
	if err != nil {
		fmt.Println("error reading stdin:", err.Error())
		return optional.Empty[command.Command]()
	}
	return textToCommand(strings.TrimSpace(text))
}
