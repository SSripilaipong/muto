package global

import (
	"slices"

	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func NewBaseModule() module.Base {
	builder := mutation.NewRuleBuilder()

	buildAll := slc.Map(builder.BuildNamedUnit)
	active := buildAll(st.FilterActiveRuleFromStatement(rawStatements))
	normal := slices.Concat(
		buildAll(st.FilterRuleFromStatement(rawStatements)),
		newForeignNormalMutators(),
	)

	collection := mutator.NewCollectionFromMutators(normal, active)
	return module.NewBase(collection, builder)
}
