package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
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

	castBoolean := func(x tokenizer.Token) stResult.Node {
		return st.NewBoolean(x.Value())
	}
	castString := func(x tokenizer.Token) stResult.Node {
		s := x.Value()
		return st.NewString(s[1 : len(s)-1])
	}
	castNumber := func(x tokenizer.Token) stResult.Node {
		return st.NewNumber(x.Value())
	}
	castClass := func(x tokenizer.Token) stResult.Node {
		return st.NewClass(x.Value())
	}
	castVariable := func(v tokenizer.Token) stResult.Node {
		return st.NewVariable(v.Value())
	}
	castObject := func(x objectNode) stResult.Node {
		return stResult.NewObject(x.Head(), x.ParamPart())
	}

	return ps.Or(
		ps.Map(castBoolean, psBase.Boolean),
		ps.Map(castString, psBase.String),
		ps.Map(castNumber, psBase.Number),
		ps.Map(castClass, psBase.ClassIncludingSymbols),
		ps.Map(castVariable, psBase.Variable),
		ps.Map(castObject, psBase.InParentheses(object)),
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
