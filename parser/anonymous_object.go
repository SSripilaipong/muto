package parser

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	"github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func anonymousObject(xs []tokenizer.Token) []tuple.Of2[anonymousObjectNode, []tokenizer.Token] {
	return ps.Map(mergeAnonymousObject, ps.Filter(filterAnonymousObject, ps.Sequence2(anonymousObjectHead, objectParamPart)))(xs)
}

var mergeAnonymousObject = tuple.Fn2(func(head st.RuleResult, params st.ObjectParamPart) anonymousObjectNode {
	return anonymousObjectNode{head: head, paramPart: params}
})

var filterAnonymousObject = tuple.Fn2(func(head st.RuleResult, param st.ObjectParamPart) bool {
	hasChildren := st.IsObjectParamPartTypeFixed(param) && st.UnsafeObjectParamPartToObjectFixedParamPart(param).Size() > 0
	return st.IsRuleResultTypeVariable(head) || hasChildren
})

func anonymousObjectHead(xs []tokenizer.Token) []tuple.Of2[st.RuleResult, []tokenizer.Token] {
	return ps.Or(
		ps.Map(mergeAnonymousObjectHeadForBoolean, boolean),
		ps.Map(mergeAnonymousObjectHeadForString, string_),
		ps.Map(mergeAnonymousObjectHeadForNumber, number),
		ps.Map(mergeAnonymousObjectHeadForVariable, variable),
		ps.Map(mergeAnonymousObjectHeadForNamedObject, inParentheses(namedObject)),
		ps.Map(mergeAnonymousObjectHeadForAnonymousObject, inParentheses(anonymousObject)),
	)(xs)
}

func mergeAnonymousObjectHeadForBoolean(x tokenizer.Token) st.RuleResult {
	return st.NewBoolean(x.Value())
}

func mergeAnonymousObjectHeadForString(x tokenizer.Token) st.RuleResult {
	s := x.Value()
	return st.NewString(s[1 : len(s)-1])
}

func mergeAnonymousObjectHeadForNumber(x tokenizer.Token) st.RuleResult {
	return st.NewNumber(x.Value())
}

func mergeAnonymousObjectHeadForVariable(v tokenizer.Token) st.RuleResult {
	return st.NewVariable(v.Value())
}

func mergeAnonymousObjectHeadForNamedObject(x namedObjectNode) st.RuleResult {
	return st.NewRuleResultNamedObject(x.Name(), x.Params())
}

func mergeAnonymousObjectHeadForAnonymousObject(x anonymousObjectNode) st.RuleResult {
	return st.NewRuleResultAnonymousObject(x.Head(), x.ParamPart())
}

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
