package base

import (
	"fmt"
	"slices"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
)

type ParamChain interface {
	DirectParams() []Node
	TotalNodes() int
	Size() int
	ReplaceChild(i int, j int, n Node) optional.Of[ParamChain]
	AppendChildrenMostOuter(children []Node) ParamChain
	All() [][]Node
	Chain(params ParamChain) ParamChain
	Clone() ParamChain
	Append(nodes []Node) ParamChain
	AppendAll(params ParamChain) ParamChain
	SliceFromOrEmpty(i int) ParamChain
	SliceUntilOrEmpty(i int) ParamChain
}

type paramChain struct {
	chain [][]Node
}

func NewParamChain(chain [][]Node) ParamChain {
	return &paramChain{chain: chain}
}

func (c *paramChain) DirectParams() []Node {
	if len(c.chain) == 0 {
		return nil
	}
	return c.chain[0]
}

func (c *paramChain) TotalNodes() int {
	var total int
	for _, nodes := range c.chain {
		total += len(nodes)
	}
	return total
}

func (c *paramChain) Size() int {
	return len(c.chain)
}

func (c *paramChain) ReplaceChild(i int, j int, n Node) optional.Of[ParamChain] {
	if !slc.ValidIndex(c.chain, i) {
		return optional.Empty[ParamChain]()
	}
	target := slices.Clone(c.chain[i])

	if !slc.ValidIndex(target, j) {
		return optional.Empty[ParamChain]()
	}
	target[j] = n

	return optional.Value(NewParamChain(append(append(c.chain[:i], target), c.chain[i+1:]...)))
}

func (c *paramChain) AppendChildrenMostOuter(children []Node) ParamChain {
	if len(c.chain) == 0 {
		return NewParamChain([][]Node{children})
	}

	i := len(c.chain) - 1
	last := append(slices.Clone(c.chain[i]), children...)
	return NewParamChain(append(c.chain[:i], last))
}

func (c *paramChain) All() [][]Node {
	return slices.Clone(c.chain)
}

func (c *paramChain) Chain(params ParamChain) ParamChain {
	if params.Size() == 0 {
		return c
	}
	paramsChain := params.All()
	if c.Size() == 0 {
		return params
	}
	left := c.chain[:slc.LastIndex(c.chain)]
	joint := append(slices.Clone(c.chain[slc.LastIndex(c.chain)]), paramsChain[0]...)
	right := paramsChain[1:]
	return NewParamChain(append(append(left, joint), right...))
}

func (c *paramChain) Clone() ParamChain {
	return NewParamChain(slices.Clone(c.chain))
}

func (c *paramChain) Append(nodes []Node) ParamChain {
	return NewParamChain(append(slices.Clone(c.chain), nodes))
}

func (c *paramChain) AppendAll(params ParamChain) ParamChain {
	return NewParamChain(append(slices.Clone(c.chain), params.All()...))
}

func (c *paramChain) SliceFromOrEmpty(i int) ParamChain {
	if !slc.ValidSliceFromIndex(c.chain, i) {
		return NewParamChain(nil)
	}
	return NewParamChain(slices.Clone(c.chain[i:]))
}

func (c *paramChain) SliceUntilOrEmpty(i int) ParamChain {
	if !slc.ValidSliceUntilIndex(i) {
		return NewParamChain(nil)
	}
	return NewParamChain(slices.Clone(c.chain[:i]))
}

func (c *paramChain) String() string {
	return fmt.Sprintf("ParamChain%s", slc.Map(slc.Map(nodeToString))(c.chain))
}

func nodeToString(x Node) string {
	return fmt.Sprintf("%s", x)
}
