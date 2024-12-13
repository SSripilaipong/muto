package program

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
)

func (w Wrapper) MutateNode(node base.Node) optional.Of[int] {
	result := w.program.MutateUntilTerminated(node)
	fmt.Println(result.TopLevelString())
	return optional.Empty[int]()
}
