package reader

import (
	"fmt"

	"github.com/SSripilaipong/muto/builder/repl/core/command"
	"github.com/SSripilaipong/muto/common/optional"
	replSt "github.com/SSripilaipong/muto/syntaxtree/repl"
)

func newReplCommand(st replSt.Command) optional.Of[command.Command] {
	switch {
	case replSt.IsQuitCommand(st):
		return optional.Value[command.Command](command.NewQuit())
	}
	fmt.Println("unknown command:", st)
	return optional.Empty[command.Command]()
}
