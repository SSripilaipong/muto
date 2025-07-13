package global

import (
	"github.com/SSripilaipong/muto/common/cliio"
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/mutation/rule"
	ruleBuilder "github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func NewModule(cliReader CliReader, cliPrinter CliPrinter) module.Module {
	builder := mutation.NewRuleBuilder(ruleBuilder.NewSimplifiedNodeBuilderFactory())

	buildAll := slc.Map(fn.Compose(mutator.ToNamedObjectMutator, builder.Build))
	active := buildAll(st.FilterActiveRuleFromStatement(rawStatements))
	normal := append(
		buildAll(st.FilterRuleFromStatement(rawStatements)),
		newForeignNormalMutators(cliReader, cliPrinter)...,
	)

	ruleCollection := mutator.NewRuleCollection(normal, active)
	return module.NewModule(ruleCollection, builder)
}

func NewModuleForStdio() module.Module {
	return NewModule(
		CliReaderFunc(cliio.ReadInputOneLine),
		CliPrinterFunc(cliio.PrintStringWithNewLine),
	)
}
