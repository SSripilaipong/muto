package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	"muto/syntaxtree"
)

func anonymousObject(xs []tokenizer.Token) []tuple.Of2[anonymousObjectNode, []tokenizer.Token] {
	return ps.Map(mergeAnonymousObject, ps.Sequence2(anonymousObjectHead, ps.OptionalGreedyRepeat(objectParam)))(xs)
}

var mergeAnonymousObject = tuple.Fn2(func(head syntaxtree.AnonymousObjectHead, params []syntaxtree.ObjectParam) anonymousObjectNode {
	return anonymousObjectNode{head: head, params: params}
})

func anonymousObjectHead(xs []tokenizer.Token) []tuple.Of2[syntaxtree.AnonymousObjectHead, []tokenizer.Token] {
	return ps.Or(
		ps.Map(mergeAnonymousObjectHeadForNamedObject, ps.Sequence3(openParenthesis, namedObject, closeParenthesis)),
		ps.Map(mergeAnonymousObjectHeadForAnonymousObject, ps.Sequence3(openParenthesis, anonymousObject, closeParenthesis)),
	)(xs)
}

var mergeAnonymousObjectHeadForNamedObject = tuple.Fn3(func(_ tokenizer.Token, x namedObjectNode, _ tokenizer.Token) syntaxtree.AnonymousObjectHead {
	return syntaxtree.NewRuleResultNamedObject(x.Name(), x.Params())
})

var mergeAnonymousObjectHeadForAnonymousObject = tuple.Fn3(func(_ tokenizer.Token, x anonymousObjectNode, _ tokenizer.Token) syntaxtree.AnonymousObjectHead {
	return syntaxtree.NewRuleResultAnonymousObject(x.Head(), x.Params())
})

type anonymousObjectNode struct {
	head   syntaxtree.AnonymousObjectHead
	params []syntaxtree.ObjectParam
}

func (n anonymousObjectNode) Params() []syntaxtree.ObjectParam {
	return n.params
}

func (n anonymousObjectNode) Head() syntaxtree.AnonymousObjectHead {
	return n.head
}
