package reader

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/muto/builder/repl/core/command"
	"github.com/SSripilaipong/muto/common/optional"
)

type Reader struct {
	lineReader LineReader
	errPrinter ErrorPrinter
}

type LineReader interface {
	ReadLine() (string, error)
}

type ErrorPrinter interface {
	Print(x string)
}

func New(lineReader LineReader, errPrinter ErrorPrinter) Reader {
	return Reader{lineReader: lineReader, errPrinter: errPrinter}
}

func (r Reader) Read() optional.Of[command.Command] {
	text, err := r.lineReader.ReadLine()
	if err != nil {
		r.errPrinter.Print(fmt.Sprint("error reading stdin:", err.Error()))
		return optional.Empty[command.Command]()
	}
	return TextToCommand(strings.TrimSpace(text))
}
