package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	tk "github.com/SSripilaipong/muto/parser/tokens"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func object(xs []tk.Token) []tuple.Of2[objectNode, []tk.Token] {
	return ps.Map(mergeObject, psBase.IgnoreSpaceBetween2(objectHead, objectParamPart))(xs)
}

var mergeObject = tuple.Fn2(func(head stResult.Node, params stResult.ParamPart) objectNode {
	return objectNode{head: head, paramPart: params}
})

func objectHead(xs []tk.Token) []tuple.Of2[stResult.Node, []tk.Token] {
	return ps.Or(
		nonNestedNode,
		ps.Map(castObjectNode, psBase.InParentheses(object)),
	)(xs)
}
