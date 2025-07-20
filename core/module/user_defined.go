package module

import (
	"github.com/SSripilaipong/muto/common/slc"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

func BuildUserDefinedModule(syntaxTree st.Module) Unattached[Serializable] {
	files := syntaxTree.Files()
	if len(files) > 1 {
		panic("currently only support at most 1 file")
	}
	var ss []stBase.Statement
	if len(files) == 1 {
		ss = files[0].Statements()
	}

	ruleBuilder := ruleMutation.NewRuleBuilder()

	buildAll := slc.Map(ruleBuilder.BuildNamedUnit)
	active := buildAll(st.FilterActiveRuleFromStatement(ss))
	normal := buildAll(st.FilterRuleFromStatement(ss))

	collection := mutator.NewCollectionFromMutators(normal, active)
	return NewUnattached(NewSerializable(NewBase(collection, ruleBuilder), syntaxTree))
}
