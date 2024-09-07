package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	"muto/syntaxtree"
)

func namedObject(xs []tokenizer.Token) []tuple.Of2[namedObjectNode, []tokenizer.Token] {
	pattern := ps.Sequence2(objectName, ps.OptionalGreedyRepeat(objectParam))
	r := ps.Map(mergeNamedObjectNode, pattern)(xs)
	return r
}

var mergeNamedObjectNode = tuple.Fn2(func(t tokenizer.Token, ps []syntaxtree.ObjectParam) namedObjectNode {
	return namedObjectNode{name: t.Value(), params: ps}
})

func objectName(xs []tokenizer.Token) []tuple.Of2[tokenizer.Token, []tokenizer.Token] {
	return ps.Or(
		nonCapitalIdentifier,
		symbol,
	)(xs)
}

type namedObjectNode struct {
	name   string
	params []syntaxtree.ObjectParam
}

func (obj namedObjectNode) Name() string {
	return obj.name
}

func (obj namedObjectNode) Params() []syntaxtree.ObjectParam {
	return obj.params
}
