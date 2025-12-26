package module

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

func BuildUserDefinedModuleFromBase(mod Base, syntaxTree st.Module) Unattached[Serializable] {
	files := syntaxTree.Files()
	if len(files) > 1 {
		panic("currently only support at most 1 file")
	}
	var ss []stBase.Statement
	if len(files) == 1 {
		ss = files[0].Statements()
	}
	buildAll := slc.Map(mod.builder.BuildNamedUnit)
	activeRules := buildAll(st.FilterActiveRuleFromStatement(ss))
	normalRules := buildAll(st.FilterRuleFromStatement(ss))

	collection := mutator.NewCollectionFromMutators(normalRules, activeRules)
	mod.ExtendCollection(collection)
	return NewUnattached(NewSerializable(mod, syntaxTree))
}
