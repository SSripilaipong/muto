package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func namedObject(xs []tokenizer.Token) []tuple.Of2[namedObjectNode, []tokenizer.Token] {
	return ps.Map(mergeNamedObjectNode, ps.Sequence2(objectName, objectParamPart))(xs)
}

var mergeNamedObjectNode = tuple.Fn2(func(t tokenizer.Token, paramPart stResult.ParamPart) namedObjectNode {
	return namedObjectNode{name: t.Value(), paramPart: paramPart}
})

func objectName(xs []tokenizer.Token) []tuple.Of2[tokenizer.Token, []tokenizer.Token] {
	return ps.Or(
		nonKeywordNonCapitalIdentifier,
		symbolName,
	)(xs)
}

type namedObjectNode struct {
	name      string
	paramPart stResult.ParamPart
}

func (obj namedObjectNode) Name() string {
	return obj.name
}

func (obj namedObjectNode) Params() stResult.ParamPart {
	return obj.paramPart
}
