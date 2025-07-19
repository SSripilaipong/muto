package portal

import (
	"bufio"
	"os"
	"strings"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"
)

type StdIn struct{}

func NewStdIn() StdIn {
	return StdIn{}
}

func (s StdIn) Call(x base.Node) optional.Of[base.Node] {
	if base.IsClassNode(x) && base.UnsafeNodeToClass(x).Name() == "$" {
		v, err := rslt.Fmap(func(s string) string {
			return strings.TrimRight(s, "\n")
		})(rslt.New(bufio.NewReader(os.Stdin).ReadString('\n'))).Return()

		if err != nil {
			return optional.Value[base.Node](base.NewErrorWithMessage(err.Error()))
		}
		return optional.Value[base.Node](base.NewString(v))
	}
	return optional.Empty[base.Node]()
}

var _ portal.Port = StdIn{}
