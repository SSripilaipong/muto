package httpserver

import (
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/module"
	mutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

const ModuleName = "http-server"

func NewModule() module.Base {
	builder := mutation.NewRuleBuilder()
	normal := slc.Map(builder.BuildNamedUnit)(st.FilterRuleFromStatement(rawStatements))
	collection := mutator.NewCollectionFromMutators(normal, nil)
	return module.NewBase(collection, builder)
}
