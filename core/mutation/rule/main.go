package mutation

import (
	"github.com/SSripilaipong/muto/core/mutation/rule/builder"
	ruleExtractor "github.com/SSripilaipong/muto/core/mutation/rule/extractor"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func New(rule st.Rule) mutator.NameWrapper {
	return mutator.NewNameWrapper(
		rule.PatternName(),
		mutator.NewReconstructor(ruleExtractor.New(rule.Pattern()), builder.New(rule.Result())),
	)
}
