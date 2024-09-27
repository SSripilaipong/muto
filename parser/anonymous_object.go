package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func anonymousObject(xs []tokenizer.Token) []tuple.Of2[anonymousObjectNode, []tokenizer.Token] {
	return ps.Map(mergeAnonymousObject, ps.Filter(filterAnonymousObject, ps.Sequence2(anonymousObjectHead, objectParamPart)))(xs)
}

var mergeAnonymousObject = tuple.Fn2(func(head stResult.Node, params stResult.ParamPart) anonymousObjectNode {
	return anonymousObjectNode{head: head, paramPart: params}
})

var filterAnonymousObject = tuple.Fn2(func(head stResult.Node, param stResult.ParamPart) bool {
	hasChildren := stResult.IsParamPartTypeFixed(param) && stResult.UnsafeParamPartToFixedParamPart(param).Size() > 0
	return stResult.IsNodeTypeVariable(head) || hasChildren
})

func anonymousObjectHead(xs []tokenizer.Token) []tuple.Of2[stResult.Node, []tokenizer.Token] {
	return ps.Or(
		ps.Map(mergeAnonymousObjectHeadForBoolean, boolean),
		ps.Map(mergeAnonymousObjectHeadForString, string_),
		ps.Map(mergeAnonymousObjectHeadForNumber, number),
		ps.Map(mergeAnonymousObjectHeadForVariable, variable),
		ps.Map(mergeAnonymousObjectHeadForNamedObject, inParentheses(namedObject)),
		ps.Map(mergeAnonymousObjectHeadForAnonymousObject, inParentheses(anonymousObject)),
	)(xs)
}

func mergeAnonymousObjectHeadForBoolean(x tokenizer.Token) stResult.Node {
	return st.NewBoolean(x.Value())
}

func mergeAnonymousObjectHeadForString(x tokenizer.Token) stResult.Node {
	s := x.Value()
	return st.NewString(s[1 : len(s)-1])
}

func mergeAnonymousObjectHeadForNumber(x tokenizer.Token) stResult.Node {
	return st.NewNumber(x.Value())
}

func mergeAnonymousObjectHeadForVariable(v tokenizer.Token) stResult.Node {
	return st.NewVariable(v.Value())
}

func mergeAnonymousObjectHeadForNamedObject(x namedObjectNode) stResult.Node {
	return stResult.NewNamedObject(x.Name(), x.Params())
}

func mergeAnonymousObjectHeadForAnonymousObject(x anonymousObjectNode) stResult.Node {
	return stResult.NewObject(x.Head(), x.ParamPart())
}

type anonymousObjectNode struct {
	head      stResult.Node
	paramPart stResult.ParamPart
}

func (n anonymousObjectNode) ParamPart() stResult.ParamPart {
	return n.paramPart
}

func (n anonymousObjectNode) Head() stResult.Node {
	return n.head
}
