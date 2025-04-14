package repl

import (
	"fmt"
	"os"

	"github.com/chzyer/readline"
	"github.com/urfave/cli/v2"

	replBuilder "github.com/SSripilaipong/muto/builder/repl"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "repl",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			os.Exit(loop(replBuilder.New(newConsoleReader(), newConsolePrinter())))
			return nil
		},
	}
}

type consolePrinter struct{}

func newConsolePrinter() consolePrinter {
	return consolePrinter{}
}

func (consolePrinter) Print(s string) { fmt.Println(s) }

type consoleReader struct {
	reader *readline.Instance
}

func newConsoleReader() consoleReader {
	reader, err := readline.New("µ> ")
	if err != nil {
		panic(fmt.Errorf("unexpected error while creating consoleReader: %w", err))
	}
	return consoleReader{reader: reader}
}

func (r consoleReader) ReadLine() (string, error) {
	return r.reader.Readline()
}
