package program

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

func (w Wrapper) MutateNode(node base.Node) optional.Of[int] {
	result := w.program.MutateUntilTerminated(node)
	w.print(result.TopLevelString())
	return optional.Empty[int]()
}

func (w Wrapper) print(x string) {
	w.printer.Print(x)
}
