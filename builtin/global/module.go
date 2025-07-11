package global

import (
	"fmt"

	"github.com/SSripilaipong/muto/common/cliio"
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/module"
	"github.com/SSripilaipong/muto/core/mutation/rule"
	ruleBuilder "github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	fileParser "github.com/SSripilaipong/muto/parser/file"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

func NewModule(cliReader CliReader, cliPrinter CliPrinter) module.Module {
	var statements []stBase.Statement
	for _, code := range codes {
		file, err := fileParser.ParseFileFromString(code).Return()
		if err != nil {
			panic(fmt.Errorf("cannot build global module: %w", err))
		}
		statements = append(statements, file.Statements()...)
	}
	linker := module.NewClassLinker()
	builder := mutation.NewRuleBuilder(ruleBuilder.NewSimplifiedNodeBuilderFactory(linker))
	buildAll := slc.Map(fn.Compose(mutator.ToNamedObjectMutator, builder.Build))
	active := buildAll(st.FilterActiveRuleFromStatement(statements))
	normal := buildAll(st.FilterRuleFromStatement(statements))

	foreignNormal := newForeignNormalMutators(cliReader, cliPrinter)

	ruleCollection := mutator.NewRuleCollection(append(normal, foreignNormal...), active)
	linker.LinkCollection(ruleCollection)

	return module.NewModule(ruleCollection, linker, builder)
}

func NewModuleForStdio() module.Module {
	return NewModule(
		CliReaderFunc(cliio.ReadInputOneLine),
		CliPrinterFunc(cliio.PrintStringWithNewLine),
	)
}

func NewModuleFromMutators(normal []mutator.NamedObjectMutator, active []mutator.NamedObjectMutator) module.Module {
	linker := module.NewClassLinker()
	builder := mutation.NewRuleBuilder(ruleBuilder.NewSimplifiedNodeBuilderFactory(linker))
	ruleCollection := mutator.NewRuleCollection(normal, active)
	linker.LinkCollection(ruleCollection)
	return module.NewModule(ruleCollection, linker, builder)
}
