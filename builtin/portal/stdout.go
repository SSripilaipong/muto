package portal

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"
)

type StdOut struct{}

func NewStdOut() StdOut {
	return StdOut{}
}

func (s StdOut) Call(x base.Node) optional.Of[base.Node] {
	if !base.IsStringNode(x) {
		return optional.Empty[base.Node]()
	}
	fmt.Println(base.UnsafeNodeToString(x).Value())
	return optional.Value[base.Node](base.Null())
}

var _ portal.Port = StdOut{}
