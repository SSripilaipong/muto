package parser

import (
	"strings"

	"phi-lang/common/parsing"
	"phi-lang/common/tuple"
	st "phi-lang/parser/syntaxtree"
	"phi-lang/tokenizer"
)

var fileParser = parsing.Map(st.NewFile, parsing.DrainLeading(tokenizer.IsLineBreak, statementsParser))

var statementsParser = parsing.GreedyRepeat(parsing.DrainTrailing(tokenizer.IsLineBreak, statementParser))

var statementParser = parsing.Map(st.RuleToStatement, ruleWithLineBreakParser)

var ruleWithLineBreakParser = parsing.Map(mergeRuleWithLineBreak, parsing.Sequence2(ruleParser, lineBreak))

var mergeRuleWithLineBreak = tuple.Fn2(func(a st.Rule, _ tokenizer.Token) st.Rule {
	return a
})

var ruleParser = parsing.Map(mergeRule, parsing.Sequence3(rulePatternParser, equalSign, ruleResultParser))

var mergeRule = tuple.Fn3(func(p st.RulePattern, _ tokenizer.Token, r st.RuleResult) st.Rule {
	return st.NewRule(p, r)
})

var rulePatternParser = parsing.Map(mergeRulePattern, parsing.ConsumeIf(tokenizer.IsIdentifier))

func mergeRulePattern(name tokenizer.Token) st.RulePattern {
	return st.NewRulePattern(name.Value())
}

var ruleResultParser = parsing.Map(mergeRuleResult, parsing.ConsumeIf(tokenizer.IsIdentifier))

func mergeRuleResult(name tokenizer.Token) st.RuleResult {
	return st.NewRuleResult(name.Value())
}

var lineBreak = parsing.ConsumeIf(tokenizer.IsLineBreak)
var equalSign = parsing.ConsumeIf(isEqualSign)

func isEqualSign(x tokenizer.Token) bool {
	return strings.TrimSpace(x.Value()) == "="
}
