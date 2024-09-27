package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/parser/tokenizer"
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
	return ps.Or(
		nonNestedNode,
		ps.Map(castObjectNode, psBase.InParentheses(object)),
	)(xs)
}
