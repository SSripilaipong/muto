package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	st "muto/syntaxtree"
)

func namedObject(xs []tokenizer.Token) []tuple.Of2[namedObjectNode, []tokenizer.Token] {
	return ps.Map(mergeNamedObjectNode, ps.Sequence2(objectName, objectParamPart))(xs)
}

func objectParamPart(xs []tokenizer.Token) []tuple.Of2[st.ObjectParamPart, []tokenizer.Token] {
	return ps.Or(
		ps.Map(mergeLeftVariadicParamPart, ps.Sequence2(variadicVar, ps.OptionalGreedyRepeat(objectParam))),
		ps.Map(mergeRightVariadicParamPart, ps.Sequence2(ps.GreedyRepeatAtLeastOnce(objectParam), variadicVar)),
		ps.Map(st.ObjectParamsToObjectFixedParamPart, ps.OptionalGreedyRepeat(objectParam)),
	)(xs)
}

var mergeLeftVariadicParamPart = tuple.Fn2(func(v variadicVarNode, params []st.ObjectParam) st.ObjectParamPart {
	return st.NewObjectLeftVariadicParamPart(v.Name(), st.ObjectFixedParamPart(params))
})

var mergeRightVariadicParamPart = tuple.Fn2(func(params []st.ObjectParam, v variadicVarNode) st.ObjectParamPart {
	return st.NewObjectRightVariadicParamPart(v.Name(), st.ObjectFixedParamPart(params))
})

var mergeNamedObjectNode = tuple.Fn2(func(t tokenizer.Token, paramPart st.ObjectParamPart) namedObjectNode {
	return namedObjectNode{name: t.Value(), paramPart: paramPart}
})

func objectName(xs []tokenizer.Token) []tuple.Of2[tokenizer.Token, []tokenizer.Token] {
	return ps.Or(
		nonCapitalIdentifier,
		symbol,
	)(xs)
}

type namedObjectNode struct {
	name      string
	paramPart st.ObjectParamPart
}

func (obj namedObjectNode) Name() string {
	return obj.name
}

func (obj namedObjectNode) Params() st.ObjectParamPart {
	return obj.paramPart
}
