package parser

import (
	ps "muto/common/parsing"
	"muto/common/tuple"
	"muto/parser/tokenizer"
	st "muto/syntaxtree"
)

func anonymousObject(xs []tokenizer.Token) []tuple.Of2[anonymousObjectNode, []tokenizer.Token] {
	return ps.Map(mergeAnonymousObject, ps.Sequence2(anonymousObjectHead, objectParamPart))(xs)
}

var mergeAnonymousObject = tuple.Fn2(func(head st.RuleResult, params st.ObjectParamPart) anonymousObjectNode {
	return anonymousObjectNode{head: head, paramPart: params}
})

func anonymousObjectHead(xs []tokenizer.Token) []tuple.Of2[st.RuleResult, []tokenizer.Token] {
	return ps.Or(
		ps.Map(mergeAnonymousObjectHeadForVariable, variable),
		ps.Map(mergeAnonymousObjectHeadForNamedObject, ps.Sequence3(openParenthesis, namedObject, closeParenthesis)),
		ps.Map(mergeAnonymousObjectHeadForAnonymousObject, ps.Sequence3(openParenthesis, anonymousObject, closeParenthesis)),
	)(xs)
}

func mergeAnonymousObjectHeadForVariable(v tokenizer.Token) st.RuleResult {
	return st.NewVariable(v.Value())
}

var mergeAnonymousObjectHeadForNamedObject = tuple.Fn3(func(_ tokenizer.Token, x namedObjectNode, _ tokenizer.Token) st.RuleResult {
	return st.NewRuleResultNamedObject(x.Name(), x.Params())
})

var mergeAnonymousObjectHeadForAnonymousObject = tuple.Fn3(func(_ tokenizer.Token, x anonymousObjectNode, _ tokenizer.Token) st.RuleResult {
	return st.NewRuleResultAnonymousObject(x.Head(), x.ParamPart())
})

type anonymousObjectNode struct {
	head      st.RuleResult
	paramPart st.ObjectParamPart
}

func (n anonymousObjectNode) ParamPart() st.ObjectParamPart {
	return n.paramPart
}

func (n anonymousObjectNode) Head() st.RuleResult {
	return n.head
}
