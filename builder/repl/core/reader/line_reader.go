package reader

import (
	"fmt"

	"github.com/chzyer/readline"
)

type lineReader struct {
	reader *readline.Instance
}

func newLineReader() lineReader {
	reader, err := readline.New("µ> ")
	if err != nil {
		panic(fmt.Errorf("unexpected error while creating lineReader: %w", err))
	}
	return lineReader{reader: reader}
}

func (r lineReader) ReadLine() (string, error) {
	return r.reader.Readline()
}
