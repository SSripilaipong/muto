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

var mergeLeftVariadicParamPart = tuple.Fn2(func(v variadicVarNode, params []st.ObjectParam) st.ObjectParamPart {
	return st.NewObjectLeftVariadicParamPart(v.Name(), params)
})

var mergeRightVariadicParamPart = tuple.Fn2(func(params []st.ObjectParam, v variadicVarNode) st.ObjectParamPart {
	return st.NewObjectRightVariadicParamPart(v.Name(), params)
})

var mergeNamedObjectNode = tuple.Fn2(func(t tokenizer.Token, paramPart st.ObjectParamPart) namedObjectNode {
	return namedObjectNode{name: t.Value(), paramPart: paramPart}
})

func objectName(xs []tokenizer.Token) []tuple.Of2[tokenizer.Token, []tokenizer.Token] {
	return ps.Or(
		nonCapitalIdentifier,
		symbolName,
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
