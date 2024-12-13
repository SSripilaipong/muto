package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SSripilaipong/muto/builder/repl/core/command"
	"github.com/SSripilaipong/muto/common/optional"
)

type Reader struct{}

func New() Reader {
	return Reader{}
}

func (Reader) Read() optional.Of[command.Command] {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("µ> ")
	text, _ := reader.ReadString('\n')
	return textToCommand(strings.TrimSpace(text))
}
