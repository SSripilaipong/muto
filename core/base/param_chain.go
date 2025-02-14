package base

import (
	"fmt"
	"slices"

	"github.com/SSripilaipong/muto/common/slc"
)

type ParamChain interface {
	DirectParams() []Node
	MostOuter() []Node
	TotalNodes() int
	Size() int
	AppendChildrenMostOuter(children []Node) ParamChain
	All() [][]Node
	Chain(params ParamChain) ParamChain
	Clone() ParamChain
	Append(nodes []Node) ParamChain
	AppendAll(params ParamChain) ParamChain
	Prepend(nodes []Node) ParamChain
	SliceFromOrEmpty(i int) ParamChain
	SliceFromNodeOrEmpty(i, j int) ParamChain
	SliceUntilOrEmpty(i int) ParamChain
	WithoutMostOuter() ParamChain
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

func (c *paramChain) MostOuter() []Node {
	if len(c.chain) == 0 {
		return nil
	}
	return c.chain[slc.LastIndex(c.chain)]
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
	if c.Size() == 0 {
		return params
	}
	paramsChain := params.All()
	left := c.chain[:slc.LastIndex(c.chain)]

	jointLeft := slices.Clone(c.chain[slc.LastIndex(c.chain)])
	jointRight := paramsChain[0]
	joint := append(jointLeft, jointRight...)

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

func (c *paramChain) Prepend(nodes []Node) ParamChain {
	return NewParamChain(append(slc.Pure(nodes), slices.Clone(c.chain)...))
}

func (c *paramChain) SliceFromNodeOrEmpty(i, j int) ParamChain {
	slicedFromChain := c.SliceFromOrEmpty(i)

	if !slc.ValidSliceFromIndex(slicedFromChain.All(), 1) {
		return NewParamChain(nil)
	}
	right := slicedFromChain.All()[1:]

	directParams := slicedFromChain.DirectParams()
	if !slc.ValidSliceFromIndex(directParams, j) {
		return NewParamChain(nil)
	}
	left := slc.Pure(slices.Clone(directParams))

	return NewParamChain(append(left, right...))
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

func (c *paramChain) WithoutMostOuter() ParamChain {
	if c.Size() == 0 {
		return NewParamChain(nil)
	}
	return c.SliceUntilOrEmpty(slc.LastIndex(c.chain))
}

func (c *paramChain) String() string {
	return fmt.Sprintf("ParamChain%s", slc.Map(slc.Map(nodeToString))(c.chain))
}

func nodeToString(x Node) string {
	return fmt.Sprintf("%s", x)
}
