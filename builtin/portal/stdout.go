package portal

import (
	"fmt"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"
)

type StdOut struct{}

func NewStdOut() StdOut {
	return StdOut{}
}

func (s StdOut) Call(nodes []base.Node) optional.Of[base.Node] {
	if len(nodes) != 1 {
		return optional.Empty[base.Node]()
	}
	x := nodes[0]
	if !base.IsStringNode(x) {
		return optional.Empty[base.Node]()
	}
	fmt.Println(base.UnsafeNodeToString(x).Value())
	return optional.Value[base.Node](base.Null())
}

var _ portal.Port = StdOut{}
