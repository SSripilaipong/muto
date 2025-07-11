package global

import (
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/mutation/rule"
	builder2 "github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
)

func NewBuiltinModuleForStdio() module.Module {
	return NewBuiltinModule(NewBuiltinMutatorsForStdio())
}

func NewBuiltinModule(builtins []mutator.NamedObjectMutator) module.Module {
	linker := module.NewClassLinker()
	ruleCollection := mutator.NewRuleCollection(builtins, nil)
	linker.LinkCollection(ruleCollection)
	builder := mutation.NewRuleBuilder(builder2.NewSimplifiedNodeBuilderFactory(linker))
	return module.NewModule(ruleCollection, linker, builder)
}
