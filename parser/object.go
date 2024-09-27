package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func object(xs []tokenizer.Token) []tuple.Of2[objectNode, []tokenizer.Token] {
	return ps.Map(mergeObject, ps.Filter(filterObject, ps.Sequence2(objectHead, objectParamPart)))(xs)
}

var mergeObject = tuple.Fn2(func(head stResult.Node, params stResult.ParamPart) objectNode {
	return objectNode{head: head, paramPart: params}
})

var filterObject = tuple.Fn2(func(head stResult.Node, param stResult.ParamPart) bool {
	hasChildren := stResult.IsParamPartTypeFixed(param) && stResult.UnsafeParamPartToFixedParamPart(param).Size() > 0
	return stResult.IsNodeTypeVariable(head) || hasChildren
})

func objectHead(xs []tokenizer.Token) []tuple.Of2[stResult.Node, []tokenizer.Token] {

	mergeObjectHeadForBoolean := func(x tokenizer.Token) stResult.Node {
		return st.NewBoolean(x.Value())
	}
	mergeObjectHeadForString := func(x tokenizer.Token) stResult.Node {
		s := x.Value()
		return st.NewString(s[1 : len(s)-1])
	}
	mergeObjectHeadForNumber := func(x tokenizer.Token) stResult.Node {
		return st.NewNumber(x.Value())
	}
	mergeObjectHeadForClass := func(x tokenizer.Token) stResult.Node {
		return st.NewClass(x.Value())
	}
	mergeObjectHeadForVariable := func(v tokenizer.Token) stResult.Node {
		return st.NewVariable(v.Value())
	}
	mergeObjectHeadForObject := func(x objectNode) stResult.Node {
		return stResult.NewObject(x.Head(), x.ParamPart())
	}

	return ps.Or(
		ps.Map(mergeObjectHeadForBoolean, boolean),
		ps.Map(mergeObjectHeadForString, string_),
		ps.Map(mergeObjectHeadForNumber, number),
		ps.Map(mergeObjectHeadForClass, classIncludingSymbols),
		ps.Map(mergeObjectHeadForVariable, variable),
		ps.Map(mergeObjectHeadForObject, inParentheses(object)),
	)(xs)
}

type objectNode struct {
	head      stResult.Node
	paramPart stResult.ParamPart
}

func (n objectNode) ParamPart() stResult.ParamPart {
	return n.paramPart
}

func (n objectNode) Head() stResult.Node {
	return n.head
}
